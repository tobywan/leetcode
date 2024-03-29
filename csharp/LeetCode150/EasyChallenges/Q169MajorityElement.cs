
using System.Xml.Linq;

namespace EasyChallenges
{
    public class Q169MajorityElement
    {
        //        Given an array nums of size n, return the majority element.

        //The majority element is the element that appears more
        //    than ⌊n / 2⌋ times.You may assume that the majorit
        //    element always exists in the array.

        //        Return k.
        public int MajorityElement(int[] nums)
        {
            var counts = new Dictionary<int, int>(50000);

            int result = 0;
            int maxCount = 0;
            int len = nums.Length;
            int threshold = len / 2;
            for (int i = 0; i < len; i++)
            {
                int n = nums[i];
                counts.TryGetValue(n, out int count);
                counts[n] = ++count;
                if (count > maxCount)
                {
                    maxCount = count;
                    result = n;
                    if (maxCount > threshold)
                    {
                        // found majority
                        return result;
                    }
                }
            }
            return result;
        }
    }
}
