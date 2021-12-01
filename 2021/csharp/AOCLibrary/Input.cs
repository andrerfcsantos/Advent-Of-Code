namespace AOCLibrary
{
    public class Input
    {
        public static List<string> GetLines(string contents)
        {
            string[] lines = contents.Split("\n");

            List<string> result = new();
            foreach (string line in lines)
            {
                string l = line.Trim();
                if (l != "")
                {
                    result.Add(l);
                }
            }

            return result;
        }
    }
}
