#[allow(dead_code)]
pub fn linear_search<T: PartialEq>(arr: &[T], target: T) -> Option<usize> {
    for (i, item) in arr.iter().enumerate() {
        if *item == target {
            return Some(i);
        }
    }

    return None;
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_linear_search() {
        let arr = [1, 2, 3, 4, 5, 6, 7, 8];
        assert_eq!(linear_search(&arr, 1), Some(0));
        assert_eq!(linear_search(&arr, 8), Some(7));
        assert_eq!(linear_search(&arr, 9), None);
    }
}