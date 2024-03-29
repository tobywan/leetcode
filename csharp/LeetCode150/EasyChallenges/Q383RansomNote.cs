
using Microsoft.VisualBasic;
using System.Xml.Linq;

namespace EasyChallenges
{
    public class Q383RansomNote
    {
        //Given two strings ransomNote and magazine,
        //return true if ransomNote can be constructed by
        //using the letters from magazine and false otherwise.

        //Each letter in magazine can only be used once in
        //ransomNote.

        public bool CanConstruct(string ransomNote, string magazine)
        {
            var ransomCharsNeeded = new Dictionary<char, int>();
            foreach(char c in ransomNote)
            {
                ransomCharsNeeded.TryGetValue(c, out int count);
                ransomCharsNeeded[c] = count + 1;
            }

            foreach (char c in magazine)
            {
                if (ransomCharsNeeded.TryGetValue(c, out int count)) {
                    int remaining = count - 1;
                    ransomCharsNeeded[c] = remaining;
                }
            }
            foreach(var pair in ransomCharsNeeded)
            {
                if (pair.Value > 0)
                {
                    // some chars were not supplied
                    return false;
                }
            }
            return true;
        }
    }
}
