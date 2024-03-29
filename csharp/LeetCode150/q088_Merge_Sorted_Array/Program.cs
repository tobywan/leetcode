// See https://aka.ms/new-console-template for more information
namespace q088
{

    public class Solution
    {
        public static void Main()
        {
        }
        public void Merge(int[] nums1, int m, int[] nums2, int n)
        {
            int mIdx = m - 1;
            int nIdx = n - 1;
            int resultIdx = m + n - 1;

            while (resultIdx >= 0)
            {
                bool pickM = (mIdx >= 0 && (nIdx < 0 || nums2[nIdx] < nums1[mIdx]));

                if (pickM)
                {
                    nums1[resultIdx] = nums1[mIdx];
                    mIdx--;
                    resultIdx--;
                }
                else
                {
                    nums1[resultIdx] = nums2[nIdx];
                    nIdx--;
                    resultIdx--;
                }
            }
        }
    }
}
