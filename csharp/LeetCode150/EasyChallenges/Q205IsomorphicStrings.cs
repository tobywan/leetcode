

using Microsoft.VisualBasic;

namespace EasyChallenges
{
    public class Q205IsomorphicStrings
    {
        //Given two strings s and t, determine if they are
        //isomorphic.

        //Two strings s and t are isomorphic if the characters in
        //s can be replaced to get t.

        //All occurrences of a character must be replaced with
        //another character while preserving the order of characters.
        //No two characters may map to the same character, but a
        //character may map to itself.

        public bool IsIsomorphic(string s, string t)
        {
            var lastSeenS = new Dictionary<char, int>();
            var lastSeenT = new Dictionary<char, int>();

            for (int i = 0; i < s.Length; i++)
            {
                char cs = s[i];
                char ct = t[i];

                lastSeenS.TryGetValue(cs, out int posS);
                lastSeenT.TryGetValue(ct, out int posT);

                if (posS != posT)
                {
                    return false;
                }
                lastSeenS[cs] = i + 1;
                lastSeenT[ct] = i + 1;
            }
            return true;
        }
    }
}
