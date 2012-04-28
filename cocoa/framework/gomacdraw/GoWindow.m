//
//  GoWindow.m
//  gomacdraw
//
//  Created by John Asmuth on 5/9/11.
//  Copyright 2011 Rutgers University. All rights reserved.
//

#import "GoWindow.h"
#import "ImageBuffer.h"

@implementation GoWindow

@synthesize imageView, eventWindow;

- (id)initWithCoder:(NSCoder *)aDecoder
{
    self = [super initWithCoder:aDecoder];
    if (self) {
        
    }
    
    return self;
}

- (void)setTitle:(NSString*)title
{
    [[self window] setTitle:title];
}
- (void)setSize:(CGSize)size
{
    size.height += 22;
    
    CGRect frame = [[self window] frame];
    frame.size = size;
    if ([self window] == nil) {
        fprintf(stderr, "nil window in gw\n");
    }
    [[self window] setFrame:frame display:NO];
    
    [imageView setFrame:frame];
    
    if (imageView == nil) {
        fprintf(stderr, "nil imageView in gw\n");
    }
    buffer = nil;
}

- (ImageBuffer*)buffer
{
    if (buffer == nil) {
        return [self newBuffer];
    }
    return buffer;
}

- (ImageBuffer*)newBuffer
{
    CGSize bufsize = [self size];
    bufsize.height -= 22;
    buffer = [[ImageBuffer alloc] initWithSize:bufsize];
    [imageView setImage:nil];
    return buffer;
}

- (CGSize)size
{
    return [[self window] frame].size;
}

- (void)flush
{
    
    CGImageRef cgimg = [[self buffer] image];
    CGSize size;
    size.width = CGImageGetWidth(cgimg);
    size.height = CGImageGetHeight(cgimg);
    
    CGSize wsize = [[self window] frame].size;
    wsize.height -= 22;
    
    NSImage* img = [[[NSImage alloc] autorelease] initWithCGImage:cgimg size:wsize];
    
    CGRect frame = [[self window] frame];
    frame.size = wsize;
    frame.origin = CGPointMake(0, 0);
    
    [imageView setFrame:frame];
    
    [imageView setImage:img];
}

- (void)dealloc
{
    [super dealloc];
}

@end
