glfw3 backend go.wde issues:

Open issues:
------------

2 - How to access OpenGL functions? This is not a real issue for go.wde because
    only At() is being used. But for the other draw2d functions in go.uik it 
    needs to be fully implemented.

3 - The event system of glfw3 works very different than that of go.wde
    However the callback functions might/should solve that

6 - How to properly shutdown the app? Why is this a problem? 
    Because glfw3.Main() is blocking a proper shutdown.


Solved issues:
--------------

1 - Where to place glfw3.Init()? go.wde doesn't work with Init(), 
    but glfw3.Init() must be run in the main thread. So placing it in the
    func wdeglfw3.init() might not solve it.
    I will test it in the gears app.
    Answ: Solved.
    
4 - How to access the buffers of a window?
    Answ: Don't. Let OpenGL handle the rendering of a window.

5 - How to implement glfw3.SwapBuffers? go.wde solves this (I guess) with
    CopyRGBA(), but glfw3 doesn't.
    Answ: By calling FlushImage() and let this function call SwapBuffers.
    
7 - The individual window rendering is done by calling 
    glfw.MakeContextCurrent. That means only one window is rendered at a time
    and kinda ruins the goroutine idea.
    Let's see how to implement it without modifying wdetest.
    Answ: Looks like that implementing it in Screen() is the easiest.
    The blocking is fixed with channels.

