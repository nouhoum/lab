object Solution {
  def main(args: Array[String]): Unit = {
    println("Hello Scala")

    // val nums1 = Array(1, 2, 3, 0, 0, 0)
    // val nums2 = Array(2, 5, 6)

    // merge(nums1, 3, nums2, 3)
    // println(nums1.mkString(", "))

    val nums11 = Array[Int](0)
    val nums22 = Array(1)

    merge(nums11, 0, nums22, 1)
    println(nums11.mkString(", "))
  }

  // https://leetcode.com/problems/merge-sorted-array
  // m: length of nums1
  // n: length of nums2
  // Time O(m + n) | Space O(1)
  def merge(nums1: Array[Int], m: Int, nums2: Array[Int], n: Int): Unit = {
    var i = n - 1 // nums2 length
    var j = m - 1 // nums1 first part length
    var k = m + n - 1
    printf("i: %d, j: %d, k: %d\n", i, j, k)
    println("---=======---")
    while (i >= 0) {
      // println(nums1(i))
      // println(nums2(j))
      println("=======")
      printf("i: %d, j: %d, k: %d\n", i, j, k)
      if (j >= 0 && nums1(j) > nums2(i)) {
        println("first IF")
        nums1(k) = nums1(j)
        k -= 1;
        j -= 1;
      } else {
        nums1(k) = nums2(i)
        k -= 1
        i -= 1
        printf(")))=======New i: %d\n", i)
      }

      println(nums1.mkString(", "))
      println("------")
    }
  }
}
