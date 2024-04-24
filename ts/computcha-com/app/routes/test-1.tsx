import type { MetaFunction } from "@remix-run/node";
import Button from "../button";
import { createHash, hash as blake3Hash } from "blake3";
import { Buffer } from "buffer";
import * as tf from "@tensorflow/tfjs";

type BufferFunction = (input: Buffer) => Buffer;

function nodeBlake3Hash(data: Buffer): Buffer {
  const hasher = createHash();
  hasher.update(data);
  return Buffer.from(hasher.digest());
}

class Gpupow {
  seed: Buffer;
  blake3Hash: BufferFunction;

  constructor(seed: Buffer, blake3Hash: BufferFunction) {
    this.seed = seed;
    this.blake3Hash = blake3Hash;
  }

  createHashBits(): Buffer {
    let currentHash = this.blake3Hash(this.seed);
    let hashIter = currentHash.values();
    let bits = Buffer.alloc(2048);
    for (let i = 0; i < 32; i++) {
      let byte = hashIter.next().value;
      if (byte === undefined) {
        currentHash = Buffer.from(currentHash);
        hashIter = currentHash.values();
        byte = hashIter.next().value;
      }
      for (let bit = 7; bit >= 0; bit--) {
        let value = (byte >> bit) & 1;
        bits[i * 8 + (7 - bit)] = value;
      }
    }
    return bits;
  }

  createTensorBits(): tf.Tensor {
    let bits = this.createHashBits();
    let tensor = tf.tensor1d(bits, "int32");
    // console.log(tensor.shape);
    return tensor;
  }

  createMatrixBits(size: number): tf.Tensor {
    let tensorBits = this.createTensorBits();
    let totalElements = size * size;
    let repeatTimes = Math.ceil(totalElements / 2048);
    let longTensor = tf.tile(tensorBits, [repeatTimes]); // Repeat the tensor to have at least 'totalElements' elements
    let longTensorSliced = longTensor.slice(0, totalElements); // Truncate the tensor to have exactly 'totalElements' elements
    let matrix = longTensorSliced.reshape([size, size]); // Reshape the tensor into a matrix
    return matrix;
  }

  reduceMatrixToVectorSum(matrix: tf.Tensor): tf.Tensor {
    return matrix.sum(1);
  }

  reduceMatrixToVectorMax(matrix: tf.Tensor): tf.Tensor {
    return matrix.max(1);
  }

  reduceMatrixToVectorMin(matrix: tf.Tensor): tf.Tensor {
    return matrix.min(1);
  }

  reduceMatrixToVectorRnd(matrix: tf.Tensor): tf.Tensor {
    // - gather the first row of the matrix
    // - each element of the first row is an index to gather from the
    //   corresponding column
    //
    // this works in the case of squaring a binary matrix, because the minimum
    // value is zero, and the maximum value is the number of rows
    let nCols = matrix.shape[1] as number;
    let nRows = matrix.shape[0] as number;
    let indices = matrix.slice([0, 0], [1, nCols]);
    indices = tf.clipByValue(indices, 0, nRows - 1);
    return matrix.gather(indices.flatten().toInt());
  }

  xorInChunks(buffer: Buffer): Buffer {
    // take in a buffer of arbitrary length, and xor it to itself in chunks of 256
    // bits (32 bytes)
    let chunkSize = 32;
    let chunks = [];
    for (let i = 0; i < buffer.length; i += chunkSize) {
      let chunk = buffer.subarray(i, i + chunkSize);
      chunks.push(chunk);
    }
    let result = Buffer.alloc(chunkSize);
    for (let chunk of chunks) {
      for (let i = 0; i < chunk.length; i++) {
        result[i] ^= chunk[i] as number;
      }
    }
    return result;
  }

  xorInChunksAlt(buffer: Buffer): Buffer {
    // take in a buffer of arbitrary length, and xor it to itself in chunks of 256
    // bits (32 bytes). testing reveals this is not any faster than the other
    // version.
    let chunkSize = 32;
    let result = Buffer.alloc(chunkSize);
    for (let i = 0; i < buffer.length; i += chunkSize) {
      let chunk = buffer.subarray(i, i + chunkSize);
      for (let j = 0; j < chunk.length; j++) {
        result[j] ^= chunk[j] as number;
      }
    }
    return result;
  }

  reduceMatrixBufferSync(matrix: tf.Tensor): Buffer {
    // the reason for reducing with these four operations is as follows. first
    // of all, what i would like to do is to hash the output matrix and then
    // send that back to the CPU. unfortuantely, GPUs do not really have that,
    // and tensorflow in particular does not support any hash functions. what i
    // want is to send the result of the matrix calculation back to the CPU
    // without sending the entire thing. how can i reduce the data, while also
    // being sure that the output is unique, and is very unlikely to be the same
    // for two different random inputs? instead of using a hash function, i thus
    // approximate a hash function by using four independent reduction methods.
    // the first is to sum all the elements of the matrix. the second is to take
    // the maximum of each row. the third is to take the minimum of each row.
    // the fourth is to take a random element from each row. the output is then
    // the concatenation of the four reductions. this is not a hash function,
    // but it is a way to reduce the data while also being sure that the output
    // is unique.
    //
    // there are other methods i could have added, such as standard deviation,
    // variance, or mean, but all of those do not actually add any information
    // to the four reduction methods provided here. instead, they are just
    // different ways of expressing the same information (e.g., both std dev,
    // var, and mean all require computing the sum first). the four reduction
    // methods provided here are independent of each other, and thus provide a
    // unique way to reduce the data. they are also as comprehensive as i can
    // figure to make it without having a hash function provided by tensorflow.
    //
    // one final advantage of these methods is that they are all highly
    // parallelizable (each column can be computed independently, and thus for
    // an NxN matrix, we have N independent threads), and thus are also suitable
    // to computation on a GPU.
    let reducedSum = this.reduceMatrixToVectorSum(matrix);
    let reducedMax = this.reduceMatrixToVectorMax(matrix);
    let reducedMin = this.reduceMatrixToVectorMin(matrix);
    let reducedRnd = this.reduceMatrixToVectorRnd(matrix);
    let reducedSumBuf = Buffer.from(reducedSum.dataSync());
    let reducedMaxBuf = Buffer.from(reducedMax.dataSync());
    let reducedMinBuf = Buffer.from(reducedMin.dataSync());
    let reducedRndBuf = Buffer.from(reducedRnd.dataSync());
    let reducedBuf = Buffer.concat([
      reducedSumBuf,
      reducedMaxBuf,
      reducedMinBuf,
      reducedRndBuf,
    ]);
    return reducedBuf;
  }

  async reduceMatrixBufferAsync(matrix: tf.Tensor): Promise<Buffer> {
    let reducedSum = this.reduceMatrixToVectorSum(matrix);
    let reducedMax = this.reduceMatrixToVectorMax(matrix);
    let reducedMin = this.reduceMatrixToVectorMin(matrix);
    let reducedRnd = this.reduceMatrixToVectorRnd(matrix);
    let reducedSumBuf = Buffer.from(await reducedSum.data());
    let reducedMaxBuf = Buffer.from(await reducedMax.data());
    let reducedMinBuf = Buffer.from(await reducedMin.data());
    let reducedRndBuf = Buffer.from(await reducedRnd.data());
    let reducedBuf = Buffer.concat([
      reducedSumBuf,
      reducedMaxBuf,
      reducedMinBuf,
      reducedRndBuf,
    ]);
    return reducedBuf;
  }

  reduceMatrixToHashSync(matrix: tf.Tensor): Buffer {
    let reducedBuf = this.reduceMatrixBufferSync(matrix);
    let xorBuf = this.xorInChunks(reducedBuf);
    let hash = this.blake3Hash(xorBuf);
    return hash;
  }

  async reduceMatrixToHashAsync(matrix: tf.Tensor): Promise<Buffer> {
    let reducedBuf = await this.reduceMatrixBufferAsync(matrix);
    let xorBuf = this.xorInChunks(reducedBuf);
    let hash = this.blake3Hash(xorBuf);
    return hash;
  }

  squareMatrix(matrix: tf.Tensor): tf.Tensor {
    return tf.matMul(matrix, matrix);
  }

  floatMatrix(matrix: tf.Tensor): tf.Tensor {
    // Set the precision of floating point operations to 32-bit
    tf.ENV.set("WEBGL_PACK", false);
    tf.ENV.set("WEBGL_RENDER_FLOAT32_ENABLED", true);

    // Check if 32-bit floating point textures are supported
    if (!tf.ENV.getBool("WEBGL_RENDER_FLOAT32_ENABLED")) {
      throw new Error(
        "This function requires 32-bit floating point textures, which are not supported on this system.",
      );
    }

    // Convert the integer matrix to a floating point matrix
    let floatMatrix = matrix.toFloat();
    let sizeFloat = tf.scalar(floatMatrix.shape[0]);
    // Apply an element-wise float point operation to the matrix that uses all
    // common floating point operations.
    let resultMatrix = floatMatrix
      .square()
      .div(sizeFloat)
      .add(sizeFloat)
      .sub(floatMatrix)
      .log()
      .square();
    let roundedMatrix = resultMatrix.round();
    let intMatrix = roundedMatrix.toInt();

    return intMatrix;
  }

  floatDivCubeMatrix(matrix: tf.Tensor): tf.Tensor {
    // Set the precision of floating point operations to 32-bit
    tf.ENV.set("WEBGL_PACK", false);
    tf.ENV.set("WEBGL_RENDER_FLOAT32_ENABLED", true);

    // Check if 32-bit floating point textures are supported
    if (!tf.ENV.getBool("WEBGL_RENDER_FLOAT32_ENABLED")) {
      throw new Error(
        "This function requires 32-bit floating point textures, which are not supported on this system.",
      );
    }

    // Convert the integer matrix to a floating point matrix
    let floatMatrix = matrix.toFloat();
    let max = floatMatrix.max();
    // floating point divide
    let divMatrix = floatMatrix.div(max);
    // cube
    let mulMatrix = divMatrix.mul(divMatrix).mul(divMatrix);
    // round to integer and convert back to integer
    let roundedMatrix = mulMatrix.round();
    let intMatrix = roundedMatrix.toInt();
    return intMatrix;
  }

  floatSquareDivMatrix(matrix: tf.Tensor): tf.Tensor {
    // Set the precision of floating point operations to 32-bit
    tf.ENV.set("WEBGL_PACK", false);
    tf.ENV.set("WEBGL_RENDER_FLOAT32_ENABLED", true);

    // Check if 32-bit floating point textures are supported
    if (!tf.ENV.getBool("WEBGL_RENDER_FLOAT32_ENABLED")) {
      throw new Error(
        "This function requires 32-bit floating point textures, which are not supported on this system.",
      );
    }

    // Convert the integer matrix to a floating point matrix
    let floatMatrix = matrix.toFloat();
    let sizeFloat = tf.scalar(floatMatrix.shape[0]);
    // floating point square and divide
    let resultMatrix = tf.matMul(floatMatrix, floatMatrix);
    let divMatrix = resultMatrix.div(sizeFloat);
    // round to integer and convert back to integer
    let roundedMatrix = divMatrix.round();
    let intMatrix = roundedMatrix.toInt();

    return intMatrix;
  }

  hashToMatrixToSquaredToReducedToHashSync(size: number): Buffer {
    let matrix = this.createMatrixBits(size);
    let squared = this.squareMatrix(matrix);
    return this.reduceMatrixToHashSync(squared);
  }

  hashToMatrixToSquaredToFloatedToReducedToHashSync(size: number): Buffer {
    let matrix = this.createMatrixBits(size);
    let squared = this.squareMatrix(matrix);
    let floated = this.floatMatrix(squared);
    return this.reduceMatrixToHashSync(floated);
  }

  hashToMatrixToSquaredToFloatSquaredToReducedToHashSync(size: number): Buffer {
    let matrix = this.createMatrixBits(size);
    let squared = this.squareMatrix(matrix);
    let floated = this.floatSquareDivMatrix(squared);
    return this.reduceMatrixToHashSync(floated);
  }

  hashToMatrixToSquaredToReducedToHashAsync(size: number): Promise<Buffer> {
    let matrix = this.createMatrixBits(size);
    let squared = this.squareMatrix(matrix);
    return this.reduceMatrixToHashAsync(squared);
  }

  hashToMatrixToSquaredToFloatDivCubeToReducedToHashSync(size: number): Buffer {
    let matrix = this.createMatrixBits(size);
    let squared = this.squareMatrix(matrix);
    let floated = this.floatDivCubeMatrix(squared);
    return this.reduceMatrixToHashSync(floated);
  }

  hashToMatrixToSquaredToFloatDivCubeToReducedToHashAsync(
    size: number,
  ): Promise<Buffer> {
    let matrix = this.createMatrixBits(size);
    let squared = this.squareMatrix(matrix);
    let floated = this.floatDivCubeMatrix(squared);
    return this.reduceMatrixToHashAsync(floated);
  }

  hashToMatrixToSquaredToFloatedToReducedToHashAsync(
    size: number,
  ): Promise<Buffer> {
    let matrix = this.createMatrixBits(size);
    let squared = this.squareMatrix(matrix);
    let floated = this.floatMatrix(squared);
    return this.reduceMatrixToHashAsync(floated);
  }

  hashToMatrixToSquaredToFloatSquaredToReducedToHashAsync(
    size: number,
  ): Promise<Buffer> {
    let matrix = this.createMatrixBits(size);
    let squared = this.squareMatrix(matrix);
    let floated = this.floatSquareDivMatrix(squared);
    return this.reduceMatrixToHashAsync(floated);
  }

  // seed -> hash -> bits -> matrix -> square -> reduce -> hash
  //
  // all performed with a matrix whose size is 1289, which is the largest prime
  // number whose cube fits into int32. the reason why the cube matters is that
  // first we square the matrix, whose max value is 1289^2, but then we also sum
  // each column in the reduction phase, meaning the true max is a cube. to do
  // this with a larger number, you would have to use int64, which is not
  // currently supported by tensorflow.
  //
  // -> hashBitMatSquareReduceHash1289
  // -> hbmsrh1289
  // -> int1289
  int1289a(): Buffer {
    return this.hashToMatrixToSquaredToReducedToHashSync(1289);
  }

  async int1289aAsync(): Promise<Buffer> {
    return this.hashToMatrixToSquaredToReducedToHashAsync(1289);
  }

  // seed -> hash -> bits -> matrix -> square -> float -> reduce -> hash
  //
  // same as above, but also with a round of element-wise floating point
  // operations which are expected to be deterministic.
  int1289b(): Buffer {
    return this.hashToMatrixToSquaredToFloatedToReducedToHashSync(1289);
  }

  async int1289bAsync(): Promise<Buffer> {
    return this.hashToMatrixToSquaredToFloatedToReducedToHashAsync(1289);
  }

  int1289c(): Buffer {
    return this.hashToMatrixToSquaredToFloatDivCubeToReducedToHashSync(1289);
  }

  async int1289cAsync(): Promise<Buffer> {
    return this.hashToMatrixToSquaredToFloatDivCubeToReducedToHashAsync(1289);
  }

  // seed -> hash -> bits -> matrix -> square -> float square -> reduce -> hash
  floatrisky(): Buffer {
    return this.hashToMatrixToSquaredToFloatSquaredToReducedToHashSync(1289);
  }

  async floatriskyAsync(): Promise<Buffer> {
    return this.hashToMatrixToSquaredToFloatSquaredToReducedToHashAsync(1289);
  }
}

export const meta: MetaFunction = () => {
  return [
    { title: "Computcha" },
    { name: "description", content: "Welcome to Computcha!" },
  ];
};

export default function Landing() {
  let blake3Hash: BufferFunction;
  if (typeof document === "undefined") {
    // running in a server environment
    blake3Hash = nodeBlake3Hash;
  } else {
    // running in a browser environment
    import("blake3/browser").then(async ({ createHash, hash }) => {
      let browserBlake3Hash = (data: Buffer) => {
        const hasher = createHash();
        hasher.update(data);
        return Buffer.from(hasher.digest());
      };
      blake3Hash = browserBlake3Hash;
    });
  }

  async function onProcessing() {
    console.log("begin");
    // gpupow int1289
    {
      let seed = Buffer.from("seed");
      let gpupow = new Gpupow(seed, blake3Hash);
      console.time("int1289");
      let promises: Promise<Buffer>[] = [];
      for (let i = 0; i < 200; i++) {
        promises.push(gpupow.int1289aAsync());
      }
      await Promise.all(promises);
      console.timeEnd("int1289");
    }
    // gpupow float1289
    {
      let seed = Buffer.from("seed");
      let gpupow = new Gpupow(seed, blake3Hash);
      console.time("float1289");
      let promises: Promise<Buffer>[] = [];
      for (let i = 0; i < 200; i++) {
        promises.push(gpupow.int1289bAsync());
      }
      await Promise.all(promises);
      console.timeEnd("float1289");
    }
    // gpupow int1289c
    {
      let seed = Buffer.from("seed");
      let gpupow = new Gpupow(seed, blake3Hash);
      console.time("int1289c");
      let promises: Promise<Buffer>[] = [];
      for (let i = 0; i < 200; i++) {
        promises.push(gpupow.int1289cAsync());
      }
      await Promise.all(promises);
      console.timeEnd("int1289c");
    }
    // gpupow floatrisky
    {
      let seed = Buffer.from("seed");
      let gpupow = new Gpupow(seed, blake3Hash);
      console.time("floatrisky");
      let promises: Promise<Buffer>[] = [];
      for (let i = 0; i < 200; i++) {
        promises.push(gpupow.floatriskyAsync());
      }
      await Promise.all(promises);
      console.timeEnd("floatrisky");
    }
    console.log("end");
  }
  return (
    <div className="">
      <div className="mb-4 mt-4 flex">
        <div className="mx-auto">
          <div className="inline-block align-middle">
            <img
              src="/earthbucks-coin.png"
              alt=""
              className="mx-auto mb-4 block h-[200px] w-[200px] rounded-full bg-[#6d3206] shadow-lg shadow-[#6d3206]"
            />
            <div className="hidden dark:block">
              <img
                src="/earthbucks-text-white.png"
                alt="Computcha"
                className="mx-auto block h-[50px]"
              />
            </div>
            <div className="block dark:hidden">
              <img
                src="/earthbucks-text-black.png"
                alt="Computcha"
                className="mx-auto block h-[50px]"
              />
            </div>
          </div>
        </div>
      </div>
      <div className="mb-4 mt-4 text-center text-black dark:text-white">
        42 trillion EBX. No pre-mine. GPUs. Big blocks. Script.
        <br />
        <br />
        Take the math test to register or log in.
      </div>
      <div className="mb-4 mt-4 h-[80px]">
        <div className="mx-auto w-[320px]">
          <Button initialText="Compute" onProcessing={onProcessing} />
        </div>
      </div>
    </div>
  );
}