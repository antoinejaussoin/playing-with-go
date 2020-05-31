using System;

namespace playing_with_go
{
  class Program
  {
    static void Main(string[] args)
    {
      Console.WriteLine("Hello World!");
      var watch = System.Diagnostics.Stopwatch.StartNew();
      var result = CalculatePi(1000000000);
      watch.Stop();
      Console.WriteLine("result " + result);
      Console.WriteLine("took " + watch.ElapsedMilliseconds + " ms");
    }

    static float CalculatePi(int numberOfIterations)
    {
      float pi = 1;
      for (int i = 1; i <= numberOfIterations; i++)
      {
        var isNegative = i % 2 == 0 ? false : true;
        int factor = i * 2 + 1;
        if (isNegative)
        {
          pi -= 1 / (float)factor;
        }
        else
        {
          pi += 1 / (float)factor;
        }
      }
      return pi * 4;
    }
  }
}


// function calculatePi(numberOfIterations)
// {
//   // Pi/4 = 1 - 1/3 + 1/5 - 1/7 + ...
//   let pi = 1;
//   for (i = 1; i <= numberOfIterations; i++)
//   {
//     const isNegative = i % 2 == 0 ? false : true;
//     const factor = i * 2 + 1
//     if (isNegative)
//     {
//       pi -= 1 / factor
//     }
//     else
//     {
//       pi += 1 / factor
//     }
//   }
//   return pi * 4
// }