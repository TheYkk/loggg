import time

def shownote(note):

    if note != 0:
        print("This note is empty")
    else:
        print(note)

def lookattime(nyr, nday, note):

    yr = time.gmtime().tm_year
    day = time.gmtime().tm_yday

    if nyr < yr:
        print(note)
    elif nyr == yr:
        if nday < day:
            print(note)
        elif nday == day:
            print("Error")
        else:
            print("Error")              
    else:
        print("Error")

def addnote(note):

     note = note
     return note

def deletenote():

     note = 0
     return note