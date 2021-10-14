# Role-based access control

- special user: `root`
- special role: `root`

## 보기

```
etcdctl user list
etcdctl role list
```

## 생성

```
etcdctl user add myuser
etcdctl role add myrole
```

## 삭제

```
etcdctl user del myuser
etcdctl role del myrole
```

## role 권한 부여

```
etcdctl role grant-permission myrole read /config/
etcdctl role grant-permission myrole write /config/
etcdctl role grant-permission myrole --prefix true readwrite /config/
```

## role 에 사용자 추가/삭제

```
etcdctl user grant-role myuser myrole
etcdctl user revoke-role myuser myrole
```

## 계정 사용 활성화

```
etcdctl user add root
etcdctl auth enable
```

## 계정 활성화 후 사용자 인증

```
etcdctl --user myuser:mypassword get sample
```

[참고문서](https://etcd.io/docs/v3.5/op-guide/authentication/)