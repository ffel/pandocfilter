Simple Math
===========

Pandoc accepts math, for instance $\frac{a}{b}$. But these fractions are
quite inconvenient to write. We can use filters to pre-process the
input.

We can write faster (or concentrate more on the writing proces) if we
can just write fraction $a/b$, a complicated fraction $$a+b / c-d$$
which is not the same as $$a + b/c - d.e$$ However, the filter should
leave a/b (not in math mode) alone.

This filter takes the character sequences before and after the / and
uses these as arguments to `\frac{}{}`. spaces immediately before and
after the / are allowed (and ignored).

For the feature: it should be possible to define patterns in the pandoc
yaml header.

One last shot to demonstrate the expansion of the multiplication dot:

$$a.b/b.d . e/f . g$$

and

$$a . b/c.d^2 . e$$

This readme can be processed as follows

    pandoc readme.md -o readme.tex --filter ./smath

