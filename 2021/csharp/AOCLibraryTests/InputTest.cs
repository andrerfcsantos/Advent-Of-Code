using Xunit;
using AOCLibrary;
using System.Collections.Generic;

namespace AOCLibraryTests
{
    public class InputTest
    {
        [Fact]
        public void TestSingleLine()
        {
            List<string> lines = Input.GetLines(@"..\..\..\testdata\single_line.txt");
            Assert.Equal(lines, new List<string> { "hello" });
        }

        [Fact]
        public void TestSingleLineNoLFEnding()
        {
            List<string> lines = Input.GetLines(@"..\..\..\testdata\single_line_nolfending.txt");
            Assert.Equal(lines, new List<string> { "hello" });
        }

        [Fact]
        public void TestTwoLines()
        {
            List<string> lines = Input.GetLines(@"..\..\..\testdata\twolines.txt");
            Assert.Equal(lines, new List<string> { "hello", "world" });
        }

        [Fact]
        public void TestTwoLinesWithEmptyLines()
        {
            List<string> lines = Input.GetLines(@"..\..\..\testdata\twolines_withempty.txt");
            Assert.Equal(lines, new List<string> { "hello", "world" });
        }
    }
}