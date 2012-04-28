//
//  main.m
//  gmdTest
//
//  Created by John Asmuth on 5/6/11.
//  Copyright 2011 Rutgers University. All rights reserved.
//

#import <Cocoa/Cocoa.h>
#import <gomacdraw/gmd.h>

int main(int argc, char *argv[])
{
    initMacDraw();
    
    GMDWindow w = openWindow();
    setWindowTitle(w, "go test");
    setWindowSize(w, 100, 100);
    showWindow(w);
    GMDImage sb = getWindowScreen(w);
    
    UInt8* data = malloc(4*100*100);
    
    for (int x=0; x<100; x++) {
        for (int y=0; y<100; y++) {
            int r = 0;
            if (x < 50) {
                r = 255;
            }
            int g = 0;
            if (y < 50) {
                g = 255;
            }
            int b = 0;
            if (y < 25 || y >= 75) {
                b = 255;
            }
            int index = 4*(x+100*y);
            data[index+0] = r;
            data[index+1] = g;
            data[index+2] = b;
            data[index+3] = 255;
            
        }
    }
    
    setScreenData(sb, data);
    
    flushWindowScreen(w);
    
    NSAppRun();
    
    releaseMacDraw();
}
