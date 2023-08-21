import file
import glob
import os
import re
import sys

if len(sys.argv) == 1:
    print("missing dir arg")
    exit()

os.chdir(sys.argv[1])
grammars = glob.glob('*.g4')
rigs = glob.glob('Test*.java')
inputs = glob.glob('*-input')
outputs = glob.glob('*-output')

for g in grammars:
    print("processing", g)
    os.system("java - Xmx500M \
              - cp .: / usr/local/lib/antlr4-complete.jar: \
              $CLASSPATH org.antlr.v4.Tool " + g)
for rig in rigs:
    match = re.search("Test([A-Za-z]+)(_[A-Za-z0-9]+)?.java", rig)
    g = match.group(1)
    testName = match.group(1)
    if match.group(2) is not None:
        testName += match.group(2)
    # print(rig, ':', g, testName)
    input = testName+'-input'
    output = testName+'-output'
    if not os.path.exists(input):
        continue
    print("# TEST", testName)
    cmd = "java -cp .:/usr/local/lib/antlr4-complete.jar:$CLASSPATH \
        Test"+testName +\
        ' < '+input +\
        ' &> /tmp/stdout'
    # print cmd
    os.system(cmd)
    expected = file(output).read()
    results = file('/tmp/stdout').read()
    if results.strip() != expected.strip():
        print("$ "+cmd)
        print("### unexpected output:")
        for l1 in results.strip().split('\n'):
            print('> '+l1)
        print('expected:')
        for l1 in expected.strip().split('\n'):
            print('> '+l1)
        print("from input ")
        for l1 in file(input).read().strip().split('\n'):
            print('< '+l1)
        print('---------')
    # for input in inputs:
    # os.system("java -cp /usr/local/lib/antlr4-complete.jar:$CLASSPATH \
    #            Test"+g+input)
