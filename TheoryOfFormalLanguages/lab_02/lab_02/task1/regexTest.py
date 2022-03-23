import re
import time

TESTS_NUM = 9
TEST_PATH = './tests/'

regex1 = re.compile(r'[1-9][.][0-9]*[E][-+][1-9][0-9]*')
regex2 = re.compile(r'[1-9][.][0-9]*?[E][-+][1-9]+?[0-9]*')
regex3 = re.compile(r'[^a-zA-Z,.!\s\";%+-][.][^a-zA-Z,.!\s\";%+-]*[E][+-][^a-zA-Z,.!\s\";%+-]+')

reg_expressions = [regex1, regex2, regex3]

test_files = [f'{TEST_PATH}test_0{i}' for i in range(1, TESTS_NUM+1)]


def read_tests(tests_list):
    tests_ = []
    for file_test in tests_list:
        with open(file_test, 'r', encoding='utf-8') as f:
            test_data = f.read()

        tests_.append(test_data)
    return tests_


if __name__ == "__main__":
    tests = read_tests(test_files)
    for regex in reg_expressions:
        print(f'Testing regex: {regex.__str__()}')
        
        for i, test in enumerate(tests):
            print(f'Test_case:{i}')
            start_time = time.perf_counter()
            regex.findall(test)
            end_time = time.perf_counter()
            print(f'Elapsed time: {(end_time - start_time) * 1000} ms')
