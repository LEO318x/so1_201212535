#include <linux/module.h>	
#include <linux/kernel.h>	
#include <linux/init.h>		
#include <linux/proc_fs.h>
#include <asm/uaccess.h>	
#include <linux/seq_file.h>
#include <linux/hugetlb.h>
#include <linux/sched.h>
#include <linux/sched/signal.h>
#include <linux/cred.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de CPU SO1");
MODULE_AUTHOR("Mike Molina");

struct task_struct *process, *process_son;
struct list_head *son;

unsigned long ramMemory;

static int write_file(struct seq_file *file , void *v)
{
	seq_printf(file, "[");
	seq_printf(file, "{\"Nombre\":\"\",\"UID\":-1,\"PID\":-1,\"Proceso\":\"\",\"Estado\":\"\",\"VmRSS\":-1,\"Hijos\":[]}");
	for_each_process(process){
		// Encontrando los procesos padre
		if (process->mm) {
            ramMemory = get_mm_rss(process->mm) << PAGE_SHIFT;
			seq_printf(file, ",{\"Nombre\":\"\",\"UID\":%d,\"PID\":%d,\"Proceso\":\"%s\",\"Estado\":\"%ld\",\"VmRSS\":%lu,\"Hijos\":[",__kuid_val(process->cred->uid), process->pid, process->comm, process->__state,ramMemory/(1024*1024));
        } else {
            seq_printf(file, ",{\"Nombre\":\"\",\"UID\":%d,\"PID\":%d,\"Proceso\":\"%s\",\"Estado\":\"%ld\",\"VmRSS\":%lu,\"Hijos\":[",__kuid_val(process->cred->uid), process->pid, process->comm, process->__state,0);
        }
		seq_printf(file, "{\"Nombre\":\"\",\"UID\":-1,\"PID\":-1,\"Proceso\":\"\",\"Estado\":\"\",\"VmRSS\":-1,\"Hijos\":[]}");
		// Encontrando los hijos de cada proceso
		list_for_each(son,&(process->children)){
			process_son = list_entry(son,struct task_struct, sibling);
			if (process_son->mm) {
            	ramMemory = get_mm_rss(process_son->mm) << PAGE_SHIFT;
				seq_printf(file, ",{\"Nombre\":\"\",\"UID\":%d,\"PID\":%d,\"Proceso\":\"%s\",\"Estado\":\"%ld\",\"VmRSS\":%lu,\"Hijos\":[]}", __kuid_val(process_son->cred->uid), process_son->pid, process_son->comm, process_son->__state,ramMemory/(1024*1024));
			} else {
				seq_printf(file, ",{\"Nombre\":\"\",\"UID\":%d,\"PID\":%d,\"Proceso\":\"%s\",\"Estado\":\"%ld\",\"VmRSS\":%lu,\"Hijos\":[]}", __kuid_val(process_son->cred->uid), process_son->pid, process_son->comm, process_son->__state,0);
			}
		}
		seq_printf(file, "]}");
	}
	seq_printf(file, "]");
	return 0;
}

static int to_open_file(struct inode *inode,struct file *file)
{
	return single_open(file,write_file,NULL);
}

static struct proc_ops operations = 
{
	.proc_open = to_open_file,
	.proc_read = seq_read
};

static int __init cpu_201212535_init(void)
{
	proc_create("cpu_201212535",0,NULL, &operations);
	printk(KERN_INFO "Mike Leonel Molina García \n");
	return 0;
}

static void __exit cpu_201212535_exit(void)
{
	remove_proc_entry("cpu_201212535",NULL);
	printk(KERN_INFO "Primer Semestre 2023\n");
}

module_init(cpu_201212535_init);
module_exit(cpu_201212535_exit);
