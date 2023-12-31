# advent2023

Working through adventofcode.com 2023 problems.

Going to try in go this year, having written basically 0 go before.
Will also probably try to throw in some ChatGPT/Copilot experiments.


## Background

For extra fun, I'm trying in a new dev environment this year. I've
never really used VSCode, and haven't written any code in a Windows
environment since... the year 2000? So I've got a Windows laptop now
(wishing it were linux, but here we are), and I'll see what I can figure out.

Download VSCode. Per [docs](https://code.visualstudio.com/docs/sourcecontrol/github)
also download git for windows, and the github VSCode extension. Click the little
branch graph on the right, and you get an option to clone a repo. Did that,
and now here I am writing a README.

Let's see about getting set up for go... install its VSCode extension.
Set up day directory and made a go.mod file, and got a warning about go not
being installed. [Download it](https://go.dev/dl/) too. Restart. In directory
ctrl-shift-p has a `Go: Initiatlize go.mod`. Prompted me to install something,
clicked ok, unclear it was going anything for a while, but eventually showed
it installing some stuff.

I've always done a lot in the command line. I'm going to avoid that a bit more
here, and try to use the VSCode UI/tooling instead. So we'll see how that goes.
Couldn't figure out how to set git user name/email without termainl.

When I added day 2, I had to make a go.work file. There are `go work use`
commands that I ran in the terminal, which edited the `go.work` file.

### Frustations

Which is to say, things I don't understand how to do yet...

1. I've got a `helpers.go` file where most of my code goes, and a `main.go` that
   drives it. I can't figure out how to re-run the last run command I did (based on
   main.go), without leaving `helpers.go` (and clicking around a few times). Maybe
   my run configs need some work, or I just need to find the right shortcut (or write
   more unit tests :)).
2. I made an infinite loop, and couldn't figure out how to stop the process. Windows
   noticed VSCode was struggling, and offered to kill it for me. Presumably there's
   another way. Ah, there were some debugger buttons at the top of the screen, I just
   didn't notice them up near the tabs of the editor portion of the windows.