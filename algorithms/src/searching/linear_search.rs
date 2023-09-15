#[allow(dead_code)]
pub fn linear_search<T: PartialEq>(arr: &[T], target: T) -> Option<usize> {
    for (i, item) in arr.iter().enumerate() {
        if *item == target {
            return Some(i);
        }
    }

    return None;
}