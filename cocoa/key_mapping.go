/*
   Copyright 2012 the go.wde authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package gmd

func containsInt(haystack []int, needle int) bool {
  for _, v := range(haystack) {
    if needle == v {
      return true
    } 
  } 
  return false
}

var blankLetterCodes = []int{71,117,115,119,116,121,122,120,99,118,96,97,98,100,101,109,10,103,111,105,107,113,123,124,125,126,63,58,55,59,56,61,54,62,60,114}
var keyMapping = map[int]string{
  0: "a",
  11: "b",
  8: "c",
  2: "d",
  14: "e",
  3: "f",
  5: "g",
  4: "h",
  34: "i",
  38: "j",
  40: "k",
  37: "l",
  46: "m",
  45: "n",
  31: "o",
  35: "p",
  12: "q",
  15: "r",
  1: "s",
  17: "t",
  32: "u",
  9: "v",
  13: "w",
  7: "x",
  16: "y",
  6: "z",
  18: "1",
  19: "2",
  20: "3",
  21: "4",
  23: "5",
  22: "6",
  26: "7",
  28: "8",
  25: "9",
  29: "0",
  83: "1", // keypad
  84: "2",
  85: "3",
  86: "4",
  87: "5",
  88: "6",
  89: "7",
  91: "8",
  92: "9",
  82: "0",
  75: "/",
  67: "*",
  78: "-",
  69: "+",
  65: ".",
  50: "`",
  27: "-",
  24: "=",
  33: "[",
  30: "]",
  42: `\`,
  41: ";",
  39: "'",
  43: ",",
  47: ".",
  44: "/",
  36: "return",
  53: "escape",
  51: "backspace",
  71: "numlock",
  117: "delete",
  115: "home",
  119: "end",
  116: "page up",
  121: "page down",
  122: "f1",
  120: "f2",
  99 : "f3",
  118: "f4",
  96 : "f5",
  97 : "f6",
  98 : "f7",
  100: "f8",
  101: "f9",
  109: "f10",
  103: "f11",
  111: "f12",
  105: "f13",
  107: "f14",
  113: "f15",
  123: "left arrow",
  124: "right arrow",
  125: "down arrow",
  126: "up arrow",
  63: "function",
  58: "alt",
  55: "super",
  59: "control",
  56: "shift",
  61: "alt",
  54: "super",
  62: "control",
  60: "shift",
  114: "insert",
  48: "tab",
  49: "space",
}

