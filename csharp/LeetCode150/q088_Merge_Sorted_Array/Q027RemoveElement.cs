
using Microsoft.VisualBasic;
using static System.Runtime.InteropServices.JavaScript.JSType;
using System.ComponentModel;
using System.Threading.Channels;
using System.Xml.Linq;
namespace EasyChallenges
{
    public class Q027RemoveElement
    {
        //        Given an integer array nums and an integer val,
        //        remove all occurrences of val in nums in-place.
        //        The order of the elements may be changed.
        //        Then return the number of elements in nums
        //        which are not equal to val.
        //        Consider the number of elements in nums which are
        //        not equal to val be k, to get accepted, you need
        //        to do the following things:
        //        Change the array nums such that the first k
        //        elements of nums contain the elements which are
        //        not equal to val.The remaining elements of nums
        //        are not important as well as the size of nums.

        //        Return k.
        public int RemoveElement(int[] nums, int val)
        {
            int leftIdx = 0;
            int rightIdx = nums.Length - 1;
            int result = nums.Length;

            while (leftIdx <= rightIdx)
            {
                if (nums[leftIdx] == val)
                {
                    if (nums[rightIdx] != val)
                    {
                        nums[leftIdx] = nums[rightIdx];
                        leftIdx++;
                    }
                    // we either moved off an occurence on the left
                    // or moved from one on the right
                    result --;
                    rightIdx--;
                }
                else
                {
                    leftIdx++;
                } 

            }
            
            
            return result;
        }
    }
}
