#[allow(dead_code)]
pub fn selection_sort<T: PartialOrd>(arr: &mut [T]) {
    let len = arr.len();

    for i in 0..len {
        let mut min_index = i;

        for j in i + 1..len {
            if arr[j] < arr[min_index] {
                min_index = j;
            }
        }

        arr.swap(i, min_index);
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_selection_sort() {
        let mut arr = [3, 2, 4, 1];
        let expected_output = [1,2,3,4];
        selection_sort(&mut arr);
        assert_eq!(arr, expected_output);

        let mut arr = [83, 22, 90, 61, 70, 56, 99, 34, 18, 75, 28, 85, 12, 46, 9, 67, 95, 41, 64, 15];
        let expected_output = [9, 12, 15, 18, 22, 28, 34, 41, 46, 56, 61, 64, 67, 70, 75, 83, 85, 90, 95, 99];
        selection_sort(&mut arr);
        assert_eq!(arr, expected_output);

        let mut arr = [44, 76, 92, 34, 55, 10, 19, 37, 71, 27, 89, 65, 52, 3, 82, 47, 97, 6, 23, 94];
        let expected_output = [3, 6, 10, 19, 23, 27, 34, 37, 44, 47, 52, 55, 65, 71, 76, 82, 89, 92, 94, 97];
        selection_sort(&mut arr);
        assert_eq!(arr, expected_output);

        let mut arr = [80, 50, 39, 67, 24, 45, 96, 87, 17, 3, 79, 62, 90, 72, 59, 9, 84, 31, 53, 14];
        let expected_output = [3, 9, 14, 17, 24, 31, 39, 45, 50, 53, 59, 62, 67, 72, 79, 80, 84, 87, 90, 96];
        selection_sort(&mut arr);
        assert_eq!(arr, expected_output);

        let mut arr = [51, 58, 94, 6, 15, 99, 78, 33, 21, 91, 42, 76, 66, 84, 11, 88, 48, 60, 37, 1];
        let expected_output = [1, 6, 11, 15, 21, 33, 37, 42, 48, 51, 58, 60, 66, 76, 78, 84, 88, 91, 94, 99];
        selection_sort(&mut arr);
        assert_eq!(arr, expected_output);
    }
}