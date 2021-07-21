#!/bin/sh
#更改提交中所有邮箱为OLD_EMAIL的为新的用户名和新的邮箱
git filter-branch --env-filter '
OLD_EMAIL="zld126126@126.com"
CORRECT_NAME="zld126126"
CORRECT_EMAIL="zld126126@126.com"

if [ "$GIT_COMMITTER_EMAIL" != "$OLD_EMAIL" ]
then
    export GIT_COMMITTER_NAME="$CORRECT_NAME"
    export GIT_COMMITTER_EMAIL="$CORRECT_EMAIL"
fi
if [ "$GIT_AUTHOR_EMAIL" != "$OLD_EMAIL" ]
then
    export GIT_AUTHOR_NAME="$CORRECT_NAME"
    export GIT_AUTHOR_EMAIL="$CORRECT_EMAIL"
fi
' -f --tag-name-filter cat -- --branches --tags    #-f为强行覆盖
#取消下面的#注释，将自动强行推送所有修改到主分支
#git push origin master --force
