<h1 align="center">Merge Sort Source</h1>

[What It Is](#what-it-is)

## What It Is

Like **[QuickSort](https://github.com/Dentrax/Data-Structures-with-Go/tree/master/quick-sort)**, Merge Sort is a Divide and Conquer algorithm. It divides input array in two halves, calls itself for the two halves and then merges the two sorted halves. The `merge()` function is used for merging two halves. The `merge(arr, l, m, r)` is key process that assumes that `arr[l..m]` and `arr[m+1..r]` are sorted and merges the two sorted sub-arrays into one. See following C implementation for details

![Preview Thumbnail](https://upload.wikimedia.org/wikipedia/commons/c/cc/Merge-sort-example-300px.gif)

MergeSort(arr[], l,  r)
--------------------------

* If r > l
* 1. Find the middle point to divide the array into two halves:  
        middle m = (l+r)/2
* 2. Call mergeSort for first half:   
        Call mergeSort(arr, l, m)
* 3. Call mergeSort for second half:
        Call mergeSort(arr, m+1, r)
* 4. Merge the two halves sorted in step 2 and 3:
        Call merge(arr, l, m, r)


The following diagram from **[Wikipedia](https://en.wikipedia.org/wiki/File:Merge_sort_algorithm_diagram.svg)** shows the complete merge sort process for an example array {38, 27, 43, 3, 9, 82, 10}. If we take a closer look at the diagram, we can see that the array is recursively divided in two halves till the size becomes 1. Once the size becomes 1, the merge processes comes into action and starts merging arrays back till the complete array is merged

![Preview Thumbnail](https://raw.githubusercontent.com/Dentrax/Data-Structures-with-Go/master/merge-sort/resources/merge-sort.png)

> * Input: {12, 11, 13, 5, 6, 7}
> * Output: Given array is `12 11 13 5 6 7`
> * Output: Sorted array is `5 6 7 11 12 13`

**Algorithm Complexity**

| Complexity		| Notation     |
| ----------------- |:------------:|
| `Time Complexity`	| `O(n log n)` |