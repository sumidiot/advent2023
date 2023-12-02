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