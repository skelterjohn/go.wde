//
//  gmd.h
//  gomacdraw
//
//  Created by John Asmuth on 5/9/11.
//  Copyright 2011 Rutgers University. All rights reserved.
//

enum GMDErrorCodes {
    GMDNoError = 0,
    GMDLoadNibError = -1,
};

typedef struct {
    char* title;
    int width, height;
} GMDWindowConfig;

typedef struct {
    char* appName;
} GMDConfig;

typedef struct {
    void* macData;
} GMDWindow;

int initMacDraw(GMDConfig* cfg);
void releaseMacDraw();

void appLoop();

GMDWindow* openWindow(GMDWindowConfig* cfg);
int closeWindow(GMDWindow* window);