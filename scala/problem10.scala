
import scala.math._
import util.control.Breaks._

/**
 * Problem 10: Summation of primes
 *
 * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 * Find the sum of all the primes below two million.
 *
 */
object Problem10 extends App
{
  override def main(args: Array[String]) {
    // Solution 1: sum all primes one by one
    var sum: Long = 0
    for (i <- 2 to 2000000) {
      if (isPrime(i)) sum += i
    }
    println(sum)

    // Solution 2: do a sieve and sum the primes
    val primes2 = sieveOfEratosthenes(2000000)
    println(primes2.sum)
  }

  /**
   * Returns the first dividend (>1) of a given number
   * @param number Long number to check
   */
  private def firstDividend(number: Long) : Long = {
    var bottom = 2L;
    val top = round(sqrt(number)) + 1;
    while (bottom < top && number%bottom != 0) bottom += 1; // find first dividend
    if (bottom < top) bottom
    else number // Prime number
  }

  /**
   * Check if `number` is a prime
   * @param number Long number to check
   */
  private def isPrime(number: Long) : Boolean = {
    firstDividend(number) == number
  }

  /**
   * Sieve of Eratosthenes
   * Returns a list of prime numbers following the sieve of Eratosthenes algorithm
   * http://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
   * @param number Long number to check
   */
  private def sieveOfEratosthenes(limit: Long) : List[Long] = {
    val flags = scala.collection.mutable.Set[Long]()
    val list = (2L to limit)

    breakable {
      var p = 2L
      while (p < round(sqrt(limit))+1) {
        var index = 2*p
        while (index <= limit) {
          flags.add(index)
          index += p
        }

        list.find(x => !flags.contains(x) && x > p) match {
          case Some(n) => p = n
          case None => break
        }
      }
    }

    (list filter { x => !flags.contains(x) }).toList // This seems slow, but it's faster than keeping track of the primes
  }

}
