package practise.leetcode;

public class RemoveDuplicatesFromSortedArray {
    public static int removeDuplicates(int[] nums) {
        if(nums.length == 0){
            return -1;
        }
        if(nums.length == 1){
            return 1;
        }
        print(nums,nums.length);
        int k = 1;
        int i = 0;
        for (int j = 1; j < nums.length; j++) {
            if(nums[i] != nums[j]){
                k++;
                i++;
                nums[i] = nums[j];
            }
        }
        print(nums,k);
        return k;
    }
    public static void print(int[] nums,int len){

        for (int i = 0; i < len; i++) {
            System.out.print(nums[i]+" ");
        }
        System.out.println();
    }

    public static void main(String[] args) {
        int nums[]= {1,1,2};
        int res = removeDuplicates(nums);
        System.out.println(res);
    }
}
