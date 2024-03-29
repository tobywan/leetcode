namespace EasyChallenges
{
    public class Q383RansomNoteTests
    {

        [Theory]
        [MemberData(nameof(GetData))]
        public void CorrectlyCanConstruct(string ransomNote, string magazine, bool expected)
        {
            var underTest = new Q383RansomNote();

            bool result = underTest.CanConstruct(ransomNote, magazine);

            Assert.Equal(expected, result);

        }

        public static IEnumerable<object[]> GetData()
        {
            //Example 1:

            //Input: ransomNote = "a", magazine = "b"
            //Output: false
            //Example 2:

            //Input: ransomNote = "aa", magazine = "ab"
            //Output: false
            //Example 3:

            //Input: ransomNote = "aa", magazine = "aab"
            //Output: true

            //Constraints:
            // 1 <= ransomNote.length, magazine.length <= 105
            //ransomNote and magazine consist of lowercase English letters.

            var allData = new List<object[]>
            {
            new object[] {
              "a",
              "b",
              false,
            },
            new object[] {
              "aa",
              "ab",
              false
            },
            new object[] {
              "aa",
              "aab",
              true,
            },
            new object[] {
              "amz",
              "abcdefghijjklmnopqrstuvwxyz",
              true,
            },
          };

            return allData;
        }
    }
}
