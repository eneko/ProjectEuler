
import util.control.Breaks._

/**
 * Problem 9: Special Pythagorean triplet
 *
 * A Pythagorean triplet is a set of three natural numbers, a  b  c, for which,
 * a*a + b*b = c*c
 *
 * For example, 32 + 42 = 9 + 16 = 25 = 52.
 *
 * There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 * Find the product abc.
 *
 */
object Problem9 extends App
{
  override def main(args: Array[String]) {
    breakable {
      for (a <- 1 to 332) {
        for (b <- a+1 to 498) {
          var c = 1000 - a - b
          // println(a,b,c)
          if (a*a + b*b == c*c) {
            println(a*b*c)
            break
          }
        }
      }
    }

  }

}
