//
//  GoWindow.h
//  gomacdraw
//
//  Created by John Asmuth on 5/9/11.
//  Copyright 2011 Rutgers University. All rights reserved.
//

#import <Foundation/Foundation.h>
#import "ImageBuffer.h"
#import "EventWindow.h"

@interface GoWindow : NSWindowController {
@private
    NSImageView* imageView;
    ImageBuffer* buffer;
    EventWindow* eventWindow;
}

@property (assign) IBOutlet NSImageView* imageView;
@property (assign) IBOutlet EventWindow* eventWindow;

- (void)setTitle:(NSString*)title;
- (void)setSize:(CGSize)size;
- (CGSize)size;
- (ImageBuffer*)newBuffer;
- (ImageBuffer*)buffer;
- (void)flush;

@end
