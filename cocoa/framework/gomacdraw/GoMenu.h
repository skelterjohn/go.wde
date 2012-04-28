//
//  GoMenu.h
//  gomacdraw
//
//  Created by John Asmuth on 5/17/11.
//  Copyright 2011 Rutgers University. All rights reserved.
//

#import <Foundation/Foundation.h>


@interface GoMenu : NSObject {
@private
    NSMenu* menu;
}

@property (assign) IBOutlet NSMenu* menu;

- (void)setAppName:(NSString*)name;

@end
