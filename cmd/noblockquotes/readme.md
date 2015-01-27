Remove Block Quotes
===================

Block quotes are easily entered in markdown documents comment on the
main text body.

> This is an example of a block quote

These block quotes are not always convenient to end up in the final
document.

So now, we can remove these:

    pandoc readme.md -o stripped.md --filter ./noblockquotes
    diff readme.md stripped.md

> Of course, it is possible to output to pdf, which will be without the
> comments.

> The diff should prove that the block quotes are gone
