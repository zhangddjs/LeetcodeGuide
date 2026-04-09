# Generated on 2026-04-09 10:22:36
# Daily practice file: sortalgo_20260409.py

from typing import List


# --------------------------------------------------------------

def heapsort(arr: List[int]) -> None:
    pass


# --------------------------------------------------------------

def quicksort(arr: List[int]) -> None:
    """
    The general idea is that you pick a "pivot" element from the array, then partition the array into two sub-arrays
    one with elements less than the pivot and one with elements greater than the pivot.
    Then, you recursively apply the same process to those sub-arrays
    Finally, you combine them all back together, and you end up with a sorted array.
    """
    pass


# --------------------------------------------------------------

def mergesort(arr: List[int]) -> None:
    """
    The basic idea is to split the array into two halves, recursively sort each half, and then merge the two sorted halves back together.
    Divide: If the array has more than one element, split it into two halves.
    Conquer: Recursively apply merge sort to each half until each sub-array has only one element (which is considered sorted).
    Merge: Combine the two sorted halves by comparing their elements and merging them in order, resulting in a fully sorted array.
    """
    pass


# --------------------------------------------------------------

def insertsort(arr: List[int]) -> None:
    """
    It starts from the second element and compares it to the elements before it, inserting it into its correct position in the sorted part. It continues this process until the entire array is sorted.
    """
    pass


# --------------------------------------------------------------

def bubblesort(arr: List[int]) -> None:
    """
    It works by repeatedly stepping through the list, comparing adjacent elements, and swapping them if they're in the wrong order. This process is repeated until the list is sorted. Typically, with each pass, the largest unsorted element "bubbles up" to its correct position.
    """
    pass
