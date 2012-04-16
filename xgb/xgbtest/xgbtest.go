/*
   Copyright 2012 John Asmuth

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

package main

import (
   "image"
   "image/color"
	"github.com/skelterjohn/go.wde/xgb"
	"github.com/skelterjohn/go.wde"
	"github.com/skelterjohn/go.wde/wdetest"
)

func wgen(width, height int) (w wde.Window, err error) {
   xw, err := xgb.NewWindow(width, height)
   w = xw

   s := 32
   icon := image.NewRGBA(image.Rectangle{Min:image.Point{0, 0}, Max:image.Point{s, s}})

   for x:=0; x<s; x++ {
      for y:=0; y<s; y++ {
      icon.Set(x, y, color.White)
      }
   }

   for i:=0; i<s; i++ {
      icon.Set(i, i, color.RGBA{255, 0, 0, 255})
   }

   xw.SetIcon(icon)

   return
}

func main() {
   wdetest.Run(wgen)
}
