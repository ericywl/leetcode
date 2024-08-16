// The beauty of an array is determined by the sum of beauty across all
// subarrays of a given size k.
// The beauty of a subarray B = [array[l], array[l+1], ..., array[r]] is quantified by
// counting the indices i that satisfy these conditions:
//	(1) l <= i <= r
//  (2) for every index j such that i < j <= r, array[i] > array[j]

use std::collections::VecDeque;

#[derive(Debug)]
pub struct MonotonicallyDecreasingDeque {
    inner: VecDeque<i16>,
}

impl MonotonicallyDecreasingDeque {
    pub fn new() -> Self {
        Self {
            inner: VecDeque::new(),
        }
    }

    pub fn push_back(&mut self, new_element: i16) {
        loop {
            match self.inner.back() {
                Some(last_element) => {
                    if last_element > &new_element {
                        self.inner.push_back(new_element);
                        break;
                    }
                    let _ = self.inner.pop_back();
                }
                None => {
                    self.inner.push_back(new_element);
                    break;
                }
            }
        }
    }

    pub fn pop_front(&mut self) -> Option<i16> {
        self.inner.pop_front()
    }

    pub fn len(&self) -> usize {
        self.inner.len()
    }
}

pub fn beauty(arr: &[i16], k: usize) -> usize {
    let mut deque = MonotonicallyDecreasingDeque::new();
    // Initialize the deque with first k elements (first window).
    for i in 0..k {
        deque.push_back(arr[i]);
    }
    // Add first window's beauty to total.
    let mut total_beauty = deque.len();
    for i in k..arr.len() {
        // If deque is of length k, it means that all elements in
        // previous window is monotonically decreasing, so we will have to
        // remove the first element before adding new one.
        if deque.len() == k {
            let _ = deque.pop_front();
        }
        // Add new element and check the deque length for beauty of
        // current window.
        deque.push_back(arr[i]);
        total_beauty += deque.len();
    }

    total_beauty
}

#[cfg(test)]
mod test {
    use super::beauty;

    #[test]
    fn example_1() {
        assert_eq!(beauty(&vec![3,6,2,9,4,1], 3), 8)
    }

    #[test]
    fn example_2() {
        assert_eq!(beauty(&vec![1,2,3,4,5,6], 3), 4)
    }

    #[test]
    fn example_3() {
        assert_eq!(beauty(&vec![3,2,2,4,1,-1], 4), 6)
    }
}