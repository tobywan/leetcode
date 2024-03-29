using System;
using static System.Runtime.InteropServices.JavaScript.JSType;
using System.ComponentModel;
using System.Xml;

public class Q026RemoveDuplicatesFromSortedArray
{
    // 26. Remove Duplicates from Sorted Array

    //    Given an integer array nums sorted in non-decreasing
    //        order, remove the duplicates in-place such that
    //        each unique element appears only once.The relative
    //        order of the elements should be kept the same.
    //        Then return the number of unique elements in nums.

    //Consider the number of unique elements of nums to be k,
    //        to get accepted, you need to do the following
    //        things:
    //Change the array nums such that the first k elements
    //        of nums contain the unique elements in the order
    //        they were present in nums initially.The remaining
    //            elements of nums are not important as well as
    //            the size of nums.
    //Return k.
    public int RemoveDuplicates(int[] nums)
    {

        int len = nums.Length;
        if (len == 0)
        {
            return 0;
        }

        int currVal = nums[0];
        int sortedIdx = 1;
        int seekerIdx = 1;


        while (seekerIdx < len)
        {
            if (currVal != nums[seekerIdx])
            {
                currVal = nums[seekerIdx];
                nums[sortedIdx] = currVal;
                sortedIdx++;
            }
            seekerIdx++;
        }

        return sortedIdx;
    }
}
