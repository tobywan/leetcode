namespace EasyChallenges
{
    public class Q205IsomorphicStringsTests
    {

        [Theory]
        [MemberData(nameof(GetData))]
        public void CheckIsIsomorphic(string s, string t, bool expected)
        {
            var underTest = new Q205IsomorphicStrings();

            bool result = underTest.IsIsomorphic(s, t);

            Assert.Equal(expected, result);

        }

        public static IEnumerable<object[]> GetData()
        {
            //Example 1:

            //Input: s = "egg", t = "add"
            //Output: true
            //Example 2:

            //Input: s = "foo", t = "bar"
            //Output: false
            //Example 3:

            //Input: s = "paper", t = "title"
            //Output: true



            //Constraints:

            //            1 <= s.length <= 5 * 104
            //t.length == s.length
            //s and t consist of any valid ascii character.

            var allData = new List<object[]>
            {
            new object[] {
              "egg",
              "add",
              true,
            },
            new object[] {
              "foo",
              "bar",
              false,
            },
            new object[] {
              "paper",
              "title",
              true,
            },
            new object[] {
              "123523734",
              "qwetweuer",
              true,
            },
            new object[] {
              "aa",
              "ab",
              false,
            },
          };

            return allData;
        }
    }
}
