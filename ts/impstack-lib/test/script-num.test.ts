import { describe, expect, test, beforeEach, it } from '@jest/globals'
import ScriptNum from '../src/script-num';

describe('ScriptNum', () => {
  const testCases = [
    { hex: '01', dec: '1' },
    { hex: 'ff', dec: '-1' },
    { hex: '0100', dec: '256' },
    { hex: 'ff00', dec: '-256' },
    { hex: '01000000', dec: '16777216' },
    { hex: 'ff000000', dec: '-16777216' },
    { hex: '0100000000000000', dec: '72057594037927936' },
    { hex: 'ff00000000000000', dec: '-72057594037927936' },
    { hex: '0100000000000000000000000000000000000000000000000000000000000000', dec: '452312848583266388373324160190187140051835877600158453279131187530910662656' },
    { hex: 'ff00000000000000000000000000000000000000000000000000000000000000', dec: '-452312848583266388373324160190187140051835877600158453279131187530910662656' },
  ];

  testCases.forEach(({ hex, dec }) => {
    test(`fromBuffer correctly converts ${hex} to ${dec}`, () => {
      const scriptNum = new ScriptNum();
      const buffer = Buffer.from(hex, 'hex');
      scriptNum.fromBuffer(buffer);
      expect(scriptNum.num.toString()).toBe(dec);
    });
  });


  testCases.forEach(({ hex, dec }) => {
    test(`toBuffer correctly converts ${dec} to ${hex}`, () => {
      const scriptNum = new ScriptNum();
      scriptNum.num = BigInt(dec);
      const buffer = scriptNum.toBuffer();
      expect(buffer.toString('hex')).toBe(hex);
    });
  });
});