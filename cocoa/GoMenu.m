//
//  GoMenu.m
//  gomacdraw
//
//  Created by John Asmuth on 5/17/11.
//  Copyright 2011 Rutgers University. All rights reserved.
//

#import "GoMenu.h"


@implementation GoMenu

@synthesize menu;

- (id)init
{
    self = [super init];
    if (self) {
        // Initialization code here.
    }
    
    return self;
}

- (void)dealloc
{
    [super dealloc];
}

- (void)setAppName:(NSString*)name
{
    [[[[menu itemAtIndex:0] submenu] itemWithTitle:@"Quit go"] setTitle:[NSString stringWithFormat:@"Quit %@", name]];
}

@end
