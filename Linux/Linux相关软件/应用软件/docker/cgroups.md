# Cgroups (Linux control groups)

Control groups， 通常称为cgroups，是Linux内核功能，允许将进程组织成分层组，然后可以限制和监视各类资源的使用。

cgroups接口是通过cgroupfs实现的。分组的核心是在Linux内核中实现的，凡是资源跟踪和限制在一组资源类型的子系统中实现（memory, CPU 等）。
