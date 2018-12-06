Dios.ros.md
Skip to content
 
Search or jump to…

Pull requests
Issues
Marketplace
Explore
 @oscarg933 Sign out
You are over your private repository plan limit (4 of 0). Please upgrade your plan, make private repositories public, or remove private repositories so that you are within your plan limit.
Your private repositories have been locked until this is resolved. Thanks for understanding. You can contact support with any questions.
1
0 30 MaxDesiatov/Carlos
forked from spring-media/Carlos
 Code  Pull requests 1  Projects 0  Wiki  Insights
Dios.ros.md #1
 Open	oscarg933 wants to merge 110 commits into MaxDesiatov:master from oscarg933:master
+3,569 −6,285 
 Conversation 0   Commits 110   Checks 0   Files changed 145
Conversation
Reviewers
No reviews
Assignees
No one assigned
Labels
None yet
Projects
None yet
Milestone
No milestone
Notifications
You’re receiving notifications because you authored the thread.
7 participants
@oscarg933
@vittoriom
@bkase
@phimage
@asalom
@toupper
@whoover
  Allow edits from maintainers.
 @oscarg933
 
oscarg933 commented just now
No description provided.

vittoriom added some commits on Feb 16, 2016
 @vittoriom
Closes spring-media#125 by updating the access date on the disk cache…  …
9a1982a
 @vittoriom
Changelog updated fo spring-media#125
d1e6cd6
 @vittoriom
Minor refactoring to Future composition
7f2ca23
 @vittoriom
Closes spring-media#123 by introducing a Result enum and using it in …  …
c21f594
 @vittoriom
Update to latest test dependencies
dc9fbdb
 @vittoriom
Closes spring-media#103 by deprecating APIs using closures instead of…  …
38c4f42
 @vittoriom
Adds a "Apps using Carlos" section
9daec47
 @vittoriom
- Closes spring-media#52  …
da52d42
 @vittoriom
Closes spring-media#130 by implementing a more careful memory managem…  …
1716056
 @vittoriom
Removes all the references to AsyncComputation
673c193
 @vittoriom
Fixes a retain cycle between Future and Promise
b6c5980
 @vittoriom
Renames Carlos Futures to Pied Piper
e5263c4
 @vittoriom
Removes the copy files phase from the frameworks
b3cc79a
 @vittoriom
Updates the pod specs for the 0.7 release
d08fb5d
 @vittoriom
Final touches
322d35a
 @vittoriom
Merge branch 'issues/52_86_100-FuturesExtraction'
de4fd30
 @vittoriom
Renames PiedPiper.podspec
a9c9a79
 @vittoriom
Closes spring-media#133 by using the keyword associatedtype instead o…  …
0e6efab
 @vittoriom
Uses fastlane to run tests (should work on Travis too)
590d040
 @vittoriom
Doesn't track the .bundle directory
5c916ea
 @vittoriom
30 seconds timeout for Xcode list
d6ec4c6
 @vittoriom
Uses Xcode 7.3 for Travis
be94447
 @vittoriom
Fixes Fastfile
4e65f0c
 @vittoriom
Updates CHANGELOGs with the Swift 2.2 migration
eb3b4ab
 @vittoriom
Implements Future.map
c8f8215
 @vittoriom
Finishes implementing map w/ tests
c1dc53b
 @vittoriom
Implements Future.filter, adds tests, updates the CHANGELOGs
7308e00
 @vittoriom
Closes spring-media#136 by moving some map variants to flatMap
a16e90b
 @vittoriom
Closes spring-media#140 by adding reduce to Future
f8a4077
 @vittoriom
Moves convenience init to Future removing them from Promise
b88c264
50 hidden items
Load more…
asalom and others added some commits on May 2, 2017
 @asalom
Make unsubscribeToMemoryWarnings global again
2bb34e3
 @asalom
Merge pull request spring-media#170 from WeltN24/fix/unsubscribeToMem…  …
fa3f472
 @asalom
improve CI
96d3130
 @asalom
Fixes CI
b062a17
 @asalom
remove carthage checkouts from repo
a54e39a
 @asalom
badges
4d99eb2
 @asalom
rename travis scripts folder and add missing bootstrap script
428f4b8
 @asalom
increase global timeout
0e3c4b0
 @asalom
add again nimble tweaks
4edd06a
 Vargas Casaseca, Cesar
MIgration to swift 4
dd982b5
 Vargas Casaseca, Cesar
Disable flaky tests.
c2281f6
 Vargas Casaseca, Cesar
update travis xcode version.
c0aac14
 @toupper
Merge pull request spring-media#176 from WeltN24/Swift4  …
2d4fb9e
 Artem Belenkov
added logs to debug crash
95b6702
 Artem Belenkov
added file name to log
cfe69f6
 Artem Belenkov
added info log level to cache levels
7290f1c
 Artem Belenkov
added logs and week self to compose method
b4872b2
 Artem Belenkov
swift 4.2 migration
5d427a7
 Artem Belenkov
changed md5 string method
18f2cb3
 Artem Belenkov
tests migrated to 4.2
acdf041
 Artem Belenkov
updated travis.yml
2e26374
 Artem Belenkov
switched to master branch of piedPieper
536bc1e
 Artem Belenkov
changed pied piper
6c8bbeb
 Artem Belenkov
reverted weak self changes
a3a7f7f
 @whoover
Merge pull request spring-media#178 from spring-media/swift_4.2  …
084c8f6
 @oscarg933
Add files via upload
c9f8566
 @oscarg933
Create Json.ros.md
8174e8a
 @oscarg933
Set theme jekyll-theme-time-machine
30fb738
 @oscarg933
Update issue templates  …
ad7e436
 @oscarg933
Add files via upload
a215dff
Merge state
Add more commits by pushing to the master branch on oscarg933/Carlos.

This branch has conflicts that must be resolved
Only those with write access to this repository can merge pull requests.
Conflicting files
CHANGELOG.md
Carlos/BasicCache.swift
Carlos/Carlos.swift
Carlos/Composed.swift
Carlos/ConditionedValueTransformation.swift
Carlos/DiskCacheLevel.swift
Carlos/Dispatched.swift
Carlos/KeyTransformation.swift
Carlos/MemoryCacheLevel.swift
Carlos/NSUserDefaultsCacheLevel.swift
Carlos/PoolCache.swift
Carlos/RequestCapperCache.swift
Carlos/SwitchCache.swift
Carlos/ValueTransformation.swift
README.md
@oscarg933
   
 
 
 
Leave a comment
Attach files by dragging & dropping, selecting them, or pasting from the clipboard.

 Styling with Markdown is supported
 ProTip! Add .patch or .diff to the end of URLs for Git’s plaintext views.
© 2018 GitHub, Inc.
Terms
Privacy
Security
Status
Help
Contact GitHub
Pricing
API
Training
Blog
About
Press h to open a hovercard with more details.
