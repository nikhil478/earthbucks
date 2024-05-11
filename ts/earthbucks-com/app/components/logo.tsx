import { $image } from "~/images";

export default function Logo() {
  // return (
  //   <div className="mx-auto block aspect-square w-[120px] rounded-full bg-black dark:bg-white p-[1px]">
  //     <div className="rounded-full bg-secondary-blue-500 p-[3px] shadow-lg shadow-[#04408d]">
  //       <div className="rounded-full bg-[#12b3ec] p-1 shadow-[inset_5px_5px_10px_#04408d]">
  //         <img src={$image("/earthbucks-wallet.png")} alt="" className="block" />
  //       </div>
  //     </div>
  //   </div>
  // );
  return (
    <div className="mx-auto block aspect-square w-[120px] ">
      <img src={$image("/earthbucks-wallet.png")} alt="" className="block" />
    </div>
  );
}
