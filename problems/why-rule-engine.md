# What is Rule Engine

to-learn List: pattern matching and conflict resolving algorithms

* If there are a large number of logics then, how you will search and apply them efficiently? \(Good performance.\)
* If logics are frequently changing and you generally code your logic in the application, then how you will manage or change the code that frequently? \(Avoid frequent deployment.\)
* Design the application such that, it can be easily maintained and understood by business people. \(Use by non-technical members\)
* If you have to keep your all business logic at a centralized place and separate from all the applications then, where you will keep it?



* **Rule:** _It is a set of the condition followed by the set of actions._ It represents the logic of the system. The rules are mainly represented in the **`if-then`** form. It contains mainly two parts, **condition**, and **action**. The rule is also known as _**production**_.

```text
Rule = Condition + Action
```



* **Expert System:** It is a program that uses the knowledge of a human expert to solve the problems and giving a solution. It is also known as a _**rule-based system**_ or _**production system**_.
* **Inference Engine:** It is a brain of _expert-system_ which manage a large number of rules and facts inside the expert system. Its job is picking rules and applying on data and generate a solution. We will discuss it in detail.

## Rule-Engine <a id="35fc"></a>

_It is an expert-system program, which runs the rules on the data and if any condition matches then it executes the corresponding actions._

![](../.gitbook/assets/image%20%2833%29.png)

it’s showed that we collect knowledge in the form of rules \(if-then form\) and stored them in any store. The rules could be stored in any storage like files or databases. Now inference engine picks the rules according to requirements and runs them on input data or query. If any patterns/condition matches then it performs the corresponding action and returns the result or solution.

## Inference-Engine <a id="fbe1"></a>

The inference engine is the component of the intelligent system in artificial intelligence, which applies logical rules to the knowledge base to infer new information from known facts. The first inference engine was part of the expert system. Inference engine commonly proceeds in two modes, which are:

1. **Forward chaining**
2. **Backward chaining**

Inference-Engine’s program works in three phases to execute the rule on given data.

![](https://miro.medium.com/max/902/1*FBztAbk-C-ip8B60dETTOA.png)

**Phase 1 — Match:** In this phase, the inference engine matches the facts and data against the set of rules. This process called pattern matching.

An algorithm which we can use for pattern matching are:

* Linear
* Rete _（Drools）_
* Treat
* Leaps

**Phase 2 — Resolve:** In this phase, the inference engine manages the order of conflicting rules. It resolves the conflict and gives the selected one rules. For resolving conflict it could use any of the following algorithms.

* Lex
* Recency
* MEA
* Refactor
* Priority wise

**Phase 3 — Execute:** In this phase, the inference engine simply runs the action of the selected rule on given data and return the output/result to the client.  


### Inference Methods: <a id="53bb"></a>

Rule engines generally use one of the following **inference methods** to implement an inference engine.

1. _Forward chaining_
2. _Backward chaining_

But before understanding the inference method, let’s understand the _**reasoning**_. There are two types of reasoning.

**1. Goal-Directed/Backward Reasoning:** It is working backward from the goal. Here we start from the main goal and then will go for sub-goals. So in goal-directed reasoning, if we want to achieve the main goal then we have to think that “_to achieve the main goal, what sub-goals we have to achieve.”_

**Example:** If we plan for an evening out, and for this, we plan to go for a movie, outing, and dinner. Then _**evening out** is our main goal_ and the _**movie, outing and dinner** are the sub-goals of the main goal_.

**2. Data-Driven/Forward Reasoning:** It starts with the available data and uses rules to extract more data until a goal is reached. Here we look at data and if we found some pattern then it performs respective action.

**Example:** Suppose we have to figure out the _**color of a pet named Fritz**_ with given rules and data.

Rules:

```text
1. If X croaks and X eats flies - Then X is a frog
2. If X chirps and X sings - Then X is a canary
3. If X is a frog - Then X is green
4. If X is a canary - Then X is yellow
```

Data:

```text
1. Fritz croaks
2. Fritz eats flies
```

Here using given rules and data we can extract more data like:

```text
Fritz is a frog.
Fritz is green.
```

So now let’s discuss the Inference Methods:

**Forward chaining:**

* It is an implementation of _Forward Reasoning_.
* It is _data-driven_.
* Facts assert into the working memory.
* One or more rules could be concurrently true.

**Backward chaining:**

* It is an implementation of _Backward Reasoning_.
* It is goal-driven.
* Start with the conclusion \(Goals\) and if not found then search for sub-goals.

There is one more category called _**Hybrid chaining**_. Drools use it. It is a combination of both forward and backward chaining.

## Advantages of rule-engine <a id="f695"></a>

We can consider the all above specific requirements in the given example as the advantages of the rule engine.

1. Rules are very easy to read and code by any non-technical person like business analyst, client team, etc. Here you have to focus on “What to do”, not “How to do”.
2. You store your all rules at center storage. This means you have a central place where your all business rules and logic exists. It will be a source of truth for you.
3. Logic is managed separately from core application logic so it can be managed and reused.
4. In rule-engine, we use different pattern matching and conflict resolving algorithms, which give high performance.
5. For frequently changing requirements, we can easily update rules. No code changes are required.
6. The complexity of code is more if it contains many decision points. The rule engine can handle much better it because they use a consistent representation of business rules.
7. The different applications can use the same rule-engine for the same logic. It increases reusability.

[https://medium.com/@er.rameshkatiyar/what-is-rule-engine-86ea759ad97d\#id\_token=eyJhbGciOiJSUzI1NiIsImtpZCI6IjZhOGJhNTY1MmE3MDQ0MTIxZDRmZWRhYzhmMTRkMTRjNTRlNDg5NWIiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJuYmYiOjE2MTY4MTkyMjksImF1ZCI6IjIxNjI5NjAzNTgzNC1rMWs2cWUwNjBzMnRwMmEyamFtNGxqZGNtczAwc3R0Zy5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsInN1YiI6IjExNzc3MzYxNzQ3MjQzOTcwNjc5MSIsImVtYWlsIjoieGlhb3l1ZXJqaXl1QGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJhenAiOiIyMTYyOTYwMzU4MzQtazFrNnFlMDYwczJ0cDJhMmphbTRsamRjbXMwMHN0dGcuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJuYW1lIjoiTGluZHNleSBXaGl0ZSIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS0vQU9oMTRHalJkQzktTnB2WjMwN1lrc0lmUFNPbXRWZXloOXRWeUp3MDNqdHA9czk2LWMiLCJnaXZlbl9uYW1lIjoiTGluZHNleSIsImZhbWlseV9uYW1lIjoiV2hpdGUiLCJpYXQiOjE2MTY4MTk1MjksImV4cCI6MTYxNjgyMzEyOSwianRpIjoiMGJiM2JkZDEwNTUwZGJhNjBiYmI0ZmJhZjdkNTBhZDIzZGFiYjJiNyJ9.frE\_ky18qkjKtX9Tc51sVouFFHR6lR143xoJZ9wEMLyb69pELvRP5DA75sr6cqEPKRfWxMyCp9GgbAXMO0iSxYktlsBqkEq4uUvGGWHnOFZet9Yfk5UXPduFFWpErPjbIR\_ZKWhtYzQLFUAOJBt1f9cvJqnPzj4g2nH8rcIWxUO2Jw0\_ZB-k6AE8MSxeIsX0O4Hb5d4HBMf9amIt0AK3BOzxskbdY7JQjD2DWq3ykVzZb7sXnXNoAbHqpa0Z5AVN4D8UPpXZAgbCjqJIIGaUsDrd3X0R\_HCYPNcEdpktPKDGnkJl4LjKGYaEoxo\_oRn4Aq9xSH4KJVtc8BgFM7JreQ](https://medium.com/@er.rameshkatiyar/what-is-rule-engine-86ea759ad97d#id_token=eyJhbGciOiJSUzI1NiIsImtpZCI6IjZhOGJhNTY1MmE3MDQ0MTIxZDRmZWRhYzhmMTRkMTRjNTRlNDg5NWIiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJuYmYiOjE2MTY4MTkyMjksImF1ZCI6IjIxNjI5NjAzNTgzNC1rMWs2cWUwNjBzMnRwMmEyamFtNGxqZGNtczAwc3R0Zy5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsInN1YiI6IjExNzc3MzYxNzQ3MjQzOTcwNjc5MSIsImVtYWlsIjoieGlhb3l1ZXJqaXl1QGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJhenAiOiIyMTYyOTYwMzU4MzQtazFrNnFlMDYwczJ0cDJhMmphbTRsamRjbXMwMHN0dGcuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJuYW1lIjoiTGluZHNleSBXaGl0ZSIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS0vQU9oMTRHalJkQzktTnB2WjMwN1lrc0lmUFNPbXRWZXloOXRWeUp3MDNqdHA9czk2LWMiLCJnaXZlbl9uYW1lIjoiTGluZHNleSIsImZhbWlseV9uYW1lIjoiV2hpdGUiLCJpYXQiOjE2MTY4MTk1MjksImV4cCI6MTYxNjgyMzEyOSwianRpIjoiMGJiM2JkZDEwNTUwZGJhNjBiYmI0ZmJhZjdkNTBhZDIzZGFiYjJiNyJ9.frE_ky18qkjKtX9Tc51sVouFFHR6lR143xoJZ9wEMLyb69pELvRP5DA75sr6cqEPKRfWxMyCp9GgbAXMO0iSxYktlsBqkEq4uUvGGWHnOFZet9Yfk5UXPduFFWpErPjbIR_ZKWhtYzQLFUAOJBt1f9cvJqnPzj4g2nH8rcIWxUO2Jw0_ZB-k6AE8MSxeIsX0O4Hb5d4HBMf9amIt0AK3BOzxskbdY7JQjD2DWq3ykVzZb7sXnXNoAbHqpa0Z5AVN4D8UPpXZAgbCjqJIIGaUsDrd3X0R_HCYPNcEdpktPKDGnkJl4LjKGYaEoxo_oRn4Aq9xSH4KJVtc8BgFM7JreQ)



## 规则引擎的优点 <a id="&#x89C4;&#x5219;&#x5F15;&#x64CE;&#x7684;&#x4F18;&#x70B9;"></a>

### 声明式编程 <a id="&#x58F0;&#x660E;&#x5F0F;&#x7F16;&#x7A0B;"></a>

规则可以很容易地解决困难的问题，并得到解决方案的验证。与代码不同，规则以较不复杂的语言编写; 业务分析师可以轻松阅读和验证一套规则。

### 逻辑和数据分离 <a id="&#x903B;&#x8F91;&#x548C;&#x6570;&#x636E;&#x5206;&#x79BB;"></a>

数据位于“域对象”中，业务逻辑位于“规则”中。根据项目的种类，这种分离是非常有利的。

### 速度和可扩展性 <a id="&#x901F;&#x5EA6;&#x548C;&#x53EF;&#x6269;&#x5C55;&#x6027;"></a>

写入Drools的Rete OO算法已经是一个成熟的算法。

在Drools的帮助下，您的应用程序变得非常可扩展。如果频繁更改请求，可以添加新规则，而无需修改现有规则。

### 知识集中化 <a id="&#x77E5;&#x8BC6;&#x96C6;&#x4E2D;&#x5316;"></a>

通过使用规则，您创建一个可执行的知识库（知识库）。这是商业政策的一个真理点。

理想情况下，规则是可读的，它们也可以用作文档。

## 规则引擎

可以将其视为一堆 if-then 语句。

精妙之处在于规则可以按任何顺序编写，引擎会决定何时使用对顺序有意义的任何方式来评估它们

