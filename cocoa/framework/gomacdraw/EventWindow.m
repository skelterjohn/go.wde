//
//  EventWindow.m
//  gomacdraw
//
//  Created by John Asmuth on 5/11/11.
//  Copyright 2011 Rutgers University. All rights reserved.
//

#import "EventWindow.h"
#import "GoWindow.h"

@implementation EventWindow

@synthesize eventQ, lock, gw;

- (id)initWithContentRect:(NSRect)contentRect styleMask:(NSUInteger)aStyle backing:(NSBackingStoreType)bufferingType defer:(BOOL)flag
{
    self = [super initWithContentRect:contentRect styleMask:aStyle backing:bufferingType defer:flag];
    if (self) {
        [self setLock:[[[NSConditionLock alloc] initWithCondition:0] autorelease]];
        [self setEventQ:[[[NSMutableArray alloc] initWithCapacity:10] autorelease]];
        [self setDelegate:self];
        [self setAcceptsMouseMovedEvents:YES];
    }
    
    return self;
}

- (void)dealloc
{
    [super dealloc];
}

- (void)setCallback:(void*)incallback
{
    
}

- (void)nq:(GMDEvent)e
{
    [lock lock];
    
    EventHolder* eh = [[[EventHolder alloc] initWithEvent:e] autorelease];
    [eventQ addObject:eh];
    
    if ([eventQ count] == 0) {
        [lock unlockWithCondition:0];
    } else {
        [lock unlockWithCondition:1];
    }
}

- (GMDEvent)dq
{
    [lock lockWhenCondition:1];
    
    EventHolder* eh = [eventQ objectAtIndex:0];
    GMDEvent e = [eh event];
    [eventQ removeObjectAtIndex:0];
    
    if ([eventQ count] == 0) {
        [lock unlockWithCondition:0];
    } else {
        [lock unlockWithCondition:1];
    }
    
    return e;
}
- (GMDEvent)mouseEvent:(NSEvent *)theEvent
{
    GMDEvent e;
    CGPoint loc = [theEvent locationInWindow];
    e.data[0] = (int)loc.x;
    e.data[1] = [self frame].size.height - (int)loc.y - 23; // not 22 so we allow 0
    e.data[2] = (int)[theEvent buttonNumber];
    return e;
}

- (void)mouseDown:(NSEvent *)theEvent
{
    GMDEvent e = [self mouseEvent:theEvent];
    e.kind = GMDMouseDown;
    [self nq:e];
}

- (void)mouseUp:(NSEvent *)theEvent
{
    GMDEvent e = [self mouseEvent:theEvent];
    e.kind = GMDMouseUp;
    [self nq:e];
}

- (void)mouseDragged:(NSEvent *)theEvent
{
    GMDEvent e = [self mouseEvent:theEvent];
    e.kind = GMDMouseDragged;
    [self nq:e];
}

- (void)mouseMoved:(NSEvent *)theEvent
{
    CGRect frameOrigin = [self frame];
    frameOrigin.origin = CGPointMake(0, 0);
    frameOrigin.size.height -= 22;
    if (!CGRectContainsPoint(frameOrigin, [theEvent locationInWindow])) {
        return;
    }
    GMDEvent e = [self mouseEvent:theEvent];
    e.kind = GMDMouseMoved;
    [self nq:e];
}

- (void)mouseEntered:(NSEvent *)theEvent
{
    GMDEvent e = [self mouseEvent:theEvent];
    e.kind = GMDMouseEntered;
    [self nq:e];
}

- (void)mouseExited:(NSEvent *)theEvent
{
    GMDEvent e = [self mouseEvent:theEvent];
    e.kind = GMDMouseExited;
    [self nq:e];
}

- (GMDEvent)keyEvent:(NSEvent*)theEvent
{
    NSString* chars = [theEvent characters];
    int keycode = [theEvent keyCode];
    GMDEvent e;
    e.data[0] = [chars characterAtIndex:0];
    e.data[1] = keycode;
    return e;
}

- (void)keyDown:(NSEvent*)theEvent
{
    GMDEvent e = [self keyEvent:theEvent];
    e.kind = GMDKeyDown;
    [self nq:e];
}

- (void)keyUp:(NSEvent*)theEvent
{
    GMDEvent e = [self keyEvent:theEvent];
    e.kind = GMDKeyUp;
    [self nq:e];
}

- (void)windowDidResize:(NSNotification *)notification
{
    NSView* view = [self contentView];
    [view removeTrackingRect:currentTrackingRect];
    currentTrackingRect = [view addTrackingRect:[view frame] owner:self userData:nil assumeInside:NO];
    GMDEvent e;
    e.kind = GMDResize;
    e.data[0] = [self frame].size.width;
    e.data[1] = [self frame].size.height-22;
    [self nq:e];
    [gw newBuffer];
}

- (void)windowWillClose:(NSNotification *)notification
{
    GMDEvent e;
    e.kind = GMDClose;
    [self nq:e];
}

@end
