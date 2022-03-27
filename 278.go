package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(i int) bool {
	return false
}

func firstBadVersion(n int) int {
	for i := 1; i <= n; i++ {
		if isBadVersion(i) {
			return i
		}
	}
	return n
}
