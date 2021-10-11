public class LongestCommonPrefix {
    public static void main(String[] args) {
        String[] words = {""};
        String prefix = longestCommonPrefix(words);
        System.out.println(prefix);

    }
    public static String longestCommonPrefix(String[] strs) {
        if(strs == null || strs.length == 0){
            return "";
        }

        String minimumLength = strs[0];
        int minLen = strs[0].length();
        if(minLen == 0){
            return "";
        }
        for (int i = 1; i < strs.length; i++) {
            if(minLen > strs[i].length()){
                minLen = strs[i].length();
                minimumLength = strs[i];
            }
        }


        String prefix = "";

        int left = 0;
        int right = minLen-1;

        do {
            int mid = (left+right)/2;
            String temp = minimumLength.substring(0,mid+1);
            boolean checkPrefix = prefixApplication(temp, strs);
            if(checkPrefix) {
                prefix = temp;
                left = mid + 1;
            }else {
                right = mid - 1;
            }

        }while (left<=right);
        return prefix;
    }

    private static boolean prefixApplication(String prefix, String[] strs) {
        for (int i = 0; i < strs.length; i++) {
            if(!strs[i].startsWith(prefix)){
                return false;
            }
        }
        return true;
    }
}
