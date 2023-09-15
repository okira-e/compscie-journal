#[allow(dead_code)]
pub fn bubble_sort<T: PartialOrd + Clone>(arr: &mut [T]) {
    for i in 0..arr.len() {
        for j in 0..arr.len() - 1 - i {
            if arr[j] > arr[j + 1] {
                let tmp = arr[j + 1].clone();
                arr[j + 1] = arr[j].clone();
                arr[j] = tmp;
            }
        }
    }
}
