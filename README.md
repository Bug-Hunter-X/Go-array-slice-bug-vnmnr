# Go Array Slice Bug

This example demonstrates a potential issue when working with slices of arrays in Go.  Because slices share the underlying array data, modifications to a slice will affect the original array.  This can lead to unexpected behavior if not handled carefully.

## The Bug

The code creates an array `a` and then creates a slice `b` from a portion of `a`. Modifying an element in `b` also changes the corresponding element in `a`.

## Solution

The solution demonstrates how to create a copy of the slice to avoid modifying the original array.