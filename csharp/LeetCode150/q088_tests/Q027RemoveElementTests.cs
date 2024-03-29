
using static System.Runtime.InteropServices.JavaScript.JSType;
using System.Collections.Generic;
using System.Xml.Linq;

namespace EasyChallenges
{
    public class Q027RemoveElementTests
    {

        [Theory]
        [MemberData(nameof(GetData))]
        public void DoesRemoveElement(int[] nums, int val, int[] expectedArray, int expectedReturn)
        {
            var underTest = new Q027RemoveElement();

            int k = underTest.RemoveElement(nums, val);

            int[] result = nums.Take(k).ToArray();
            Array.Sort(result);
            Assert.Equal(expectedArray,result);

            Assert.Equal(expectedReturn, k);
        }

        public static IEnumerable<object[]> GetData()
        {
            //Example 1:

            //Input: nums = [3, 2, 2, 3], val = 3
            //Output: 2, nums = [2, 2, _, _]
            //Explanation: Your function should return k = 2, with the first two elements of nums being 2.
            //It does not matter what you leave beyond the returned k(hence they are underscores).
            //Example 2:

            //Input: nums = [0, 1, 2, 2, 3, 0, 4, 2], val = 2
            //Output: 5, nums = [0, 1, 4, 0, 3, _, _, _]
            //Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
            //Note that the five elements can be returned in any order.
            //It does not matter what you leave beyond the returned k(hence they are underscores).


            //Constraints:

            //0 <= nums.length <= 100
            //0 <= nums[i] <= 50
            //0 <= val <= 100

            var allData = new List<object[]>
            {
            new object[] {
              new int[]{ 3,2,2,3 },
              3,
              new int[]{ 2,2 },
              2 ,
            },
            new object[] {
              new int[]{0, 1, 2, 2, 3, 0, 4, 2},
              2,
              new int[]{ 0,0,1,3,4 },
              5 ,
            },
            new object[] {
              new int[]{9,9,9,9,9,9,9,9,9,9},
              9,
              Array.Empty<int>(),
              0,
            },
          };

            return allData;
        }
    }
}
