namespace EasyChallenges
{
    public class Q026RemoveDuplicatesFromSortedArrayTests
    {
        [Theory]
        [MemberData(nameof(GetData))]
        public void DoesRemoveDuplicates(int[] nums, int[] expected)
        {
            Q026RemoveDuplicatesFromSortedArray uut = new Q026RemoveDuplicatesFromSortedArray();

            int result = uut.RemoveDuplicates(nums);

            Assert.Equal(nums.Take(expected.Length), expected);
            Assert.Equal(expected.Length, result);

        }

        public static IEnumerable<object[]> GetData()
        {
            //    Example 1:

            //Input: nums = [1,1,2]
            //    Output: 2, nums = [1,2, _]
            //    Explanation: Your function should return k = 2,
            //    with the first two elements of nums being 1 and 2
            //    respectively.
            //    It does not matter what you leave beyond the
            //    returned k(hence they are underscores).
            //Example 2:

            //Input: nums = [0,0,1,1,1,2,2,3,3,4]
            //    Output: 5, nums = [0,1,2,3,4, _, _, _, _, _]
            //    Explanation: Your function should return k = 5,
            //    with the first five elements of nums being
            //    0, 1, 2, 3, and 4 respectively.
            //    It does not matter what you leave beyond the
            //    returned k(hence they are underscores).

            var allData = new List<object[]>
            {
                new object[]
                {
                    new int[] { 1, 1, 2},
                    new int[] { 1, 2 },
                },
                new object[]
                {
                    new int[] {0,0,1,1,1,2,2,3,3,4},
                    new int[] {0,1,2,3,4},
                },
                new object[]
                {
                    Array.Empty<int>(),
                    Array.Empty<int>(),
                },
                new object[]
                {
                    new int[] {0,0,0,0,0,0},
                    new int[] {0},
                },
                new object[]
                {
                    new int[] {1},
                    new int[] {1},
                },

            };

            return allData;
        }

    }
}
