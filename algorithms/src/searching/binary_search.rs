#[allow(dead_code)]
pub fn binary_search<T: PartialOrd>(arr: &[T], target: T) -> Option<usize> {
    let mut low = 0;
    let mut high = arr.len() - 1;

    while low <= high {
        let mid = (low + high) / 2;
        let guess = &arr[mid];

        if *guess == target {
            return Some(mid);
        } else if *guess > target {
            high = mid - 1;
        } else {
            low = mid + 1;
        }
    }

    return None;
}
