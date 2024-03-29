namespace EasyChallenges
{
    public class Q169MajorityElementTests
    {

        [Theory]
        [MemberData(nameof(GetData))]
        public void FindsMajorityElement(int[] nums, int expected)
        {
            var underTest = new Q169MajorityElement();

            int result = underTest.MajorityElement(nums );

            Assert.Equal( expected, result );

        }

        public static IEnumerable<object[]> GetData()
        {
            //Example 1:

            //Input: nums = [3, 2, 3]
            //Output: 3

            //Example 2:

            //Input: nums = [2, 2, 1, 1, 1, 2, 2]
            //Output: 2

            //Constraints:
            //            n == nums.length
            //1 <= n <= 5 * 104
            //- 109 <= nums[i] <= 109

            //Follow - up: Could you solve the problem in linear time and in O(1) space?

            var allData = new List<object[]>
            {
            new object[] {
              new int[]{ 3,2,3 },
              3,
            },
            new object[] {
              new int[]{2, 2, 1, 1, 1, 2, 2},
              2,
            },
            new object[] {
              new int[]{1},
              1,
            },
            new object[] {
              new int[]{1,2,3,3,3,3,3,1,2},
              3,
            },
          };

            return allData;
        }
    }
}
