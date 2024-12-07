def main():
    file_content = open("input.txt", "r", encoding="utf-8").read()
    
    list = file_content.split("\n")
        
    after_num = {}
    result_list_to_check_mid = []
    
    for elem in list:
        elem_len = len(elem)
        
        if elem_len == 5:
            left = elem[:2] # It would be better to find an index of "|" but anyway :)
            right = elem[3:]
            
            if left in after_num.keys():
               list = after_num[left]
               list.append(right)
            else:
                after_num[left] = [right]
        
        elif elem_len > 5:
            elements_list = elem.split(",")
            elements = elements_list
            elements = elem.split(",")
            elements.reverse()
            
            is_intersection = False
                
            for index, item in enumerate(elements):
                
                elem_to_check = elements[index : index + 1][0]
                list_to_check_in = elements[index + 1:]
                
                elem_list = after_num[elem_to_check]                
                intersection_list = intersection(list_to_check_in, elem_list)
                
                if len(intersection_list) > 0:
                    is_intersection = True
                    break
            
            if is_intersection == False:
                result_list_to_check_mid.append(elements_list)

    
    print("result_list_to_check_mid = ", result_list_to_check_mid)        
        
    sum = 0
    if len(result_list_to_check_mid) > 0:
        for l in result_list_to_check_mid:
            elem_list_len = len(l)
            middle_index = int((elem_list_len - 1)/2)
            middle_elem = int(l[middle_index])
            sum += middle_elem
            
    print("sum = ", sum)
    return

def intersection(lst1, lst2):
    return list(set(lst1) & set(lst2))


main()