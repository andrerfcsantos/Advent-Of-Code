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

        public static List<List<string>> GetGroupedLines(string contents)
        {
            string[] lines = contents.Split("\n");

            List<List<string>> result = new();
            List<string> chunk = new();

            foreach (string line in lines)
            {
                string l = line.Trim();
                if (l != "")
                {
                    chunk.Add(l);
                } else
                {
                    result.Add(chunk);
                    chunk = new();
                }
            }

            result.Add(chunk);

            return result;
        }
    }
}
