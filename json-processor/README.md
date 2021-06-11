## Foxit - JSON Processor



Use any tool that you are comfortable with to process this JSON file into required output schema. 

- The output should be JSON lines for each block of `result` in `completions`. 
- The `label` would be `htmllabels`, `sentence` would be `text` of each result and `contract_id` would be `data.filename` and remove the suffix `.html`. 
- Please sort these JSON lines by ascending order of interval between `endOffset` and `startOffset`. (Optional)



### Required Output Schema

```json
{
    "label": "Contract Title",
    "sentence": "Mutual Confidentiality and Non-Disclosure Agreement",
    "contract_id": "1qpqQsbhhUq"
}
```

### Raw Data

```json
{
    "completions": [
        {
            "created_at": 1612850273,
            "id": 2001,
            "lead_time": 2552.251,
            "result": [
                {
                    "from_name": "ner",
                    "id": "IpAlj7cp63",
                    "to_name": "text",
                    "type": "hypertextlabels",
                    "value": {
                        "end": "/h1[1]/a[1]/text()[1]",
                        "endOffset": 24,
                        "htmllabels": [
                            "Contract Title"
                        ],
                        "start": "/h1[1]/text()[1]",
                        "startOffset": 0,
                        "text": "Mutual Confidentiality and Non-Disclosure Agreement"
                    }
                },
                {
                    "from_name": "ner",
                    "id": "GrK2W6QwCs",
                    "to_name": "text",
                    "type": "hypertextlabels",
                    "value": {
                        "end": "/p[2]/text()[1]",
                        "endOffset": 352,
                        "htmllabels": [
                            "Commencement"
                        ],
                        "start": "/p[2]/text()[1]",
                        "startOffset": 0,
                        "text": "In order to protect certain confidential information which may be disclosed between them, NewNet Communication Technologies, LLC and the undersigned company (each referred to as a “Party” or collectively as the “Parties”) agree to the following terms and conditions (the “Agreement”) to cover disclosure of the Confidential Information described below:"
                    }
                },
                {
                    "from_name": "ner",
                    "id": "gXVcsH6WEX",
                    "to_name": "text",
                    "type": "hypertextlabels",
                    "value": {
                        "end": "/p[21]/text()[2]",
                        "endOffset": 155,
                        "htmllabels": [
                            "CI Definition"
                        ],
                        "start": "/p[21]/text()[2]",
                        "startOffset": 2,
                        "text": "Either Party may terminate this Agreement by giving the other Party written notice thereof at least thirty (30) days prior to the effective date thereof."
                    }
                },
                {
                    "from_name": "ner",
                    "id": "mNB7O1NnNG",
                    "to_name": "text",
                    "type": "hypertextlabels",
                    "value": {
                        "end": "/p[21]/text()[3]",
                        "endOffset": 125,
                        "htmllabels": [
                            "Survival"
                        ],
                        "start": "/p[21]/text()[3]",
                        "startOffset": 1,
                        "text": "However, the obligations of this Agreement shall survive termination for the three (3) year period described in paragraph 3."
                    }
                },
                {
                    "from_name": "ner",
                    "id": "U4eD_TQ614",
                    "to_name": "text",
                    "type": "hypertextlabels",
                    "value": {
                        "end": "/p[24]/text()[2]",
                        "endOffset": 271,
                        "htmllabels": [
                            "Assignment and Succssor"
                        ],
                        "start": "/p[24]/text()[2]",
                        "startOffset": 3,
                        "text": "Neither Party may delegate its obligations hereunder or assign its rights as a Recipient without the prior written consent of the other Party, and any purported assignment or delegation in violation of this Agreement will be void and deemed a breach of this Agreement."
                    }
                },
                {
                    "from_name": "ner",
                    "id": "BwoMq7Bjp3",
                    "to_name": "text",
                    "type": "hypertextlabels",
                    "value": {
                        "end": "/p[25]/text()[2]",
                        "endOffset": 386,
                        "htmllabels": [
                            "Entire Agreement"
                        ],
                        "start": "/p[25]/text()[2]",
                        "startOffset": 0,
                        "text": "Except by a specific written instrument duly executed by both Paties, this Agreement sets forth the entire understanding of the Parties with respect to the subject matter hereof; incorporates and merges any and all previous agreements, understandings, and communications (oral or written) with respect to the subject matter of this Agreement and may not be modified, amended or waived. "
                    }
                }
            ]
        }
    ],
    "data": {
        "filename": "1qpqQsbhhUq.html",
        "text": "<html> html_content </html>"
    },
    "id": 23
}
```



### Note

- I Also sort these JSON lines by ascending order of interval between **startOffset** and **endOffset**.
- As you can see its result and print each interval value.
- Maybe, it's fine not looking like output schema. Show extra column **Interval**.



### How to start

- Run command as follows

```sh
go run ./src/main.go`
go build -o {app_name} ./src
./{app_name}
```


- You're ready to go.
- And you'll see it printing on system console.



### Test & Result

```json
[
  {
    "label": "Contract Title",
    "sentence": "Mutual Confidentiality and Non-Disclosure Agreement",
    "contract_id": "1qpqQsbhhUq",
    "interval": 24
  },
  {
    "label": "Survival",
    "sentence": "However, the obligations of this Agreement shall survive termination for the three (3) year period described in paragraph 3.",
    "contract_id": "1qpqQsbhhUq",
    "interval": 124
  },
  {
    "label": "CI Definition",
    "sentence": "Either Party may terminate this Agreement by giving the other Party written notice thereof at least thirty (30) days prior to the effective date thereof.",
    "contract_id": "1qpqQsbhhUq",
    "interval": 153
  },
  {
    "label": "Assignment and Succssor",
    "sentence": "Neither Party may delegate its obligations hereunder or assign its rights as a Recipient without the prior written consent of the other Party, and any purported assignment or
delegation in violation of this Agreement will be void and deemed a breach of this Agreement.",
    "contract_id": "1qpqQsbhhUq",
    "interval": 268
  },
  {
    "label": "Commencement",
    "sentence": "In order to protect certain confidential information which may be disclosed between them, NewNet Communication Technologies, LLC and the undersigned company (each referred to
as a “Party” or collectively as the “Parties”) agree to the following terms and conditions (the “Agreement”) to cover disclosure of the Confidential Information described below:",
    "contract_id": "1qpqQsbhhUq",
    "interval": 352
  },
  {
    "label": "Entire Agreement",
    "sentence": "Except by a specific written instrument duly executed by both Paties, this Agreement sets forth the entire understanding of the Parties with respect to the subject matter hereof; incorporates and merges any and all previous agreements, understandings, and communications (oral or written) with respect to the subject matter of this Agreement and may not be modified,
amended or waived. ",
    "contract_id": "1qpqQsbhhUq",
    "interval": 386
  }
]
```

