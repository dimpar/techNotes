## 7 rulses of a great Git commit message:

- Separate subject from body with a blank line
- Limit the subject line to 50 characters
- Capitalize the subject line
- Do not end the subject line with a period
- Use the imperative mood in the subject line
- Wrap the body at 72 characters
- Use the body to explain what and why vs. how

Git commit messages: 
A properly formed Git commit subject line should always be able to complete the following sentence:

If applied, this commit will your subject line here
For example:

If applied, this commit will refactor subsystem X for readability
If applied, this commit will update getting started documentation
If applied, this commit will remove deprecated methods
If applied, this commit will release version 1.0.0
If applied, this commit will merge pull request #123 from user/branch
Notice how this doesn’t work for the other non-imperative forms:

If applied, this commit will fixed bug with Y
If applied, this commit will changing behavior of X
If applied, this commit will more fixes for broken stuff
If applied, this commit will sweet new API methods

## Git useful commands:

Good looking logs
- `- git log --graph --all --oneline --decorate`

Rebasing with preserving merges (should have a clean tree with only your changes in the PR)
- git rebase master --preserve-merges (deprecated)
- git rebase master --rebase-merges

If you screwed up the rebase and have the entire changes that are already in master
- git checkout master
- git pull
- git checkout -b your_new_branch
- git cherry-pick 40sdf32d
...
- git branch -D your_old_branch
- git branch -m your_new_branch your_old_branch (rename)
- git push origin your_new_branch --force (you will loose the prev commit history)
