
import util.control.Breaks._

/**
 * Problem 9: Special Pythagorean triplet
 */
object Problem9 extends App
{
  override def main(args: Array[String]) {
    var a=1
    var b=2
    var c=3

    breakable {
      for (a <- 1 to 333) {
        for (b <- (a+1) to 500) {
          var c = 1000 - a - b
          println(a,b,c)
          if (a*a + b*b == c*c) {
            println(a*b*c)
            break
          }
        }
      }
    }

  }

}
