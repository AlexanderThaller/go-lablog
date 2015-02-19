package scm

import (
	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
	git "github.com/libgit2/git2go"
)

const (
	//Name is the name of the current package. Used in logging.
	Name = "scm"
)

//Commit will add and commit the given entry into the repository that lays unter
//the given datadir.
func Commit(datadir string, entry data.Entry) error {
	l := logger.New(Name, "Commit")

	repo, tree, parent, err := openRepository(datadir, entry.GetProject().Name+
		".csv")
	if err != nil {
		return errgo.Notef(err, "can not open repository")
	}

	message := entry.GetProject().Name + " - " + entry.Type() + " - " +
		entry.GetTimeStamp().Format(data.EntryCSVTimeStampFormat)

	sig, err := Signature()
	if err != nil {
		return errgo.Notef(err, "can not get signature")
	}

	var commitID *git.Oid
	if parent == nil {
		commitID, err = repo.CreateCommit("HEAD", sig, sig, message, tree)
	} else {
		commitID, err = repo.CreateCommit("HEAD", sig, sig, message, tree, parent)
	}
	if err != nil {
		return errgo.Notef(err, "can not create commit")
	}

	l.Trace("CommitID: ", commitID)

	return nil
}

//Signature returns the signature of the current user based on the .gitconfig
//file.
func Signature() (*git.Signature, error) {
	sig := git.Signature{
		Name:  "Alexander Thaller",
		Email: "alexander.thaller@atraveo.de",
	}

	return &sig, nil
}

func openRepository(datadir, filename string) (repo *git.Repository, tree *git.Tree, parent *git.Commit, err error) {
	repo, err = git.OpenRepository(datadir)
	if err != nil {
		err = errgo.Notef(err, "can not open repository")
		return
	}

	index, err := repo.Index()
	if err != nil {
		err = errgo.Notef(err, "can not get index of repository")
		return
	}

	err = index.AddByPath(filename)
	if err != nil {
		err = errgo.Notef(err, "can not add project file to index of repository")
		return
	}

	treeID, err := index.WriteTree()
	if err != nil {
		err = errgo.Notef(err, "can not write tree of index")
		return
	}

	currBranch, err := repo.Head()
	if err == nil {
		parent, err = repo.LookupCommit(currBranch.Target())
		if err != nil {
			err = errgo.Notef(err, "can not get current tip")
			return
		}
	}

	tree, err = repo.LookupTree(treeID)
	if err != nil {
		err = errgo.Notef(err, "can not get tree from treeID")
		return
	}

	return
}
