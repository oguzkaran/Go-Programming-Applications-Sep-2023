package org.csystem.android.app.generator.random.databinding;
import org.csystem.android.app.generator.random.R;
import org.csystem.android.app.generator.random.BR;
import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import android.view.View;
@SuppressWarnings("unchecked")
public class ActivityMainBindingImpl extends ActivityMainBinding implements org.csystem.android.app.generator.random.generated.callback.OnClickListener.Listener {

    @Nullable
    private static final androidx.databinding.ViewDataBinding.IncludedLayouts sIncludes;
    @Nullable
    private static final android.util.SparseIntArray sViewsWithIds;
    static {
        sIncludes = null;
        sViewsWithIds = null;
    }
    // views
    @NonNull
    private final android.widget.LinearLayout mboundView0;
    @NonNull
    private final android.widget.EditText mboundView1;
    @NonNull
    private final android.widget.ListView mboundView10;
    @NonNull
    private final android.widget.EditText mboundView2;
    @NonNull
    private final android.widget.EditText mboundView3;
    @NonNull
    private final android.widget.EditText mboundView4;
    @NonNull
    private final android.widget.EditText mboundView5;
    @NonNull
    private final android.widget.Button mboundView6;
    @NonNull
    private final android.widget.Button mboundView7;
    @NonNull
    private final android.widget.Button mboundView8;
    @NonNull
    private final android.widget.Button mboundView9;
    // variables
    @Nullable
    private final android.view.View.OnClickListener mCallback4;
    @Nullable
    private final android.view.View.OnClickListener mCallback2;
    @Nullable
    private final android.view.View.OnClickListener mCallback3;
    @Nullable
    private final android.view.View.OnClickListener mCallback1;
    // values
    // listeners
    // Inverse Binding Event Handlers
    private androidx.databinding.InverseBindingListener mboundView1androidTextAttrChanged = new androidx.databinding.InverseBindingListener() {
        @Override
        public void onChange() {
            // Inverse of serverInfo.host
            //         is serverInfo.setHost((java.lang.String) callbackArg_0)
            java.lang.String callbackArg_0 = androidx.databinding.adapters.TextViewBindingAdapter.getTextString(mboundView1);
            // localize variables for thread safety
            // serverInfo.host
            java.lang.String serverInfoHost = null;
            // serverInfo
            org.csystem.android.app.generator.random.viewmodel.data.ServerInfo serverInfo = mServerInfo;
            // serverInfo != null
            boolean serverInfoJavaLangObjectNull = false;



            serverInfoJavaLangObjectNull = (serverInfo) != (null);
            if (serverInfoJavaLangObjectNull) {




                serverInfo.setHost(((java.lang.String) (callbackArg_0)));
            }
        }
    };
    private androidx.databinding.InverseBindingListener mboundView2androidTextAttrChanged = new androidx.databinding.InverseBindingListener() {
        @Override
        public void onChange() {
            // Inverse of PortStringConverter.INSTANCE.toStr(serverInfo.port)
            //         is serverInfo.setPort((int) PortStringConverter.INSTANCE.toPort(callbackArg_0))
            java.lang.String callbackArg_0 = androidx.databinding.adapters.TextViewBindingAdapter.getTextString(mboundView2);
            // localize variables for thread safety
            // PortStringConverter.INSTANCE.toStr(serverInfo.port)
            java.lang.String portStringConverterINSTANCEToStrServerInfoPort = null;
            // serverInfo.port
            int serverInfoPort = 0;
            // serverInfo
            org.csystem.android.app.generator.random.viewmodel.data.ServerInfo serverInfo = mServerInfo;
            // serverInfo != null
            boolean serverInfoJavaLangObjectNull = false;



            serverInfoJavaLangObjectNull = (serverInfo) != (null);
            if (serverInfoJavaLangObjectNull) {



                if ((org.csystem.android.app.generator.random.viewmodel.converter.PortStringConverter.INSTANCE) != (null)) {



                    org.csystem.android.app.generator.random.viewmodel.converter.PortStringConverter.INSTANCE.toPort(callbackArg_0);


                    serverInfo.setPort(((int) (org.csystem.android.app.generator.random.viewmodel.converter.PortStringConverter.INSTANCE.toPort(callbackArg_0))));
                }
            }
        }
    };
    private androidx.databinding.InverseBindingListener mboundView3androidTextAttrChanged = new androidx.databinding.InverseBindingListener() {
        @Override
        public void onChange() {
            // Inverse of CountStringConverter.INSTANCE.toStr(generateInfo.count)
            //         is generateInfo.setCount((int) CountStringConverter.INSTANCE.toCount(callbackArg_0))
            java.lang.String callbackArg_0 = androidx.databinding.adapters.TextViewBindingAdapter.getTextString(mboundView3);
            // localize variables for thread safety
            // generateInfo.count
            int generateInfoCount = 0;
            // CountStringConverter.INSTANCE.toStr(generateInfo.count)
            java.lang.String countStringConverterINSTANCEToStrGenerateInfoCount = null;
            // generateInfo != null
            boolean generateInfoJavaLangObjectNull = false;
            // generateInfo
            org.csystem.android.app.generator.random.viewmodel.data.GenerateInfo generateInfo = mGenerateInfo;



            generateInfoJavaLangObjectNull = (generateInfo) != (null);
            if (generateInfoJavaLangObjectNull) {



                if ((org.csystem.android.app.generator.random.viewmodel.converter.CountStringConverter.INSTANCE) != (null)) {



                    org.csystem.android.app.generator.random.viewmodel.converter.CountStringConverter.INSTANCE.toCount(callbackArg_0);


                    generateInfo.setCount(((int) (org.csystem.android.app.generator.random.viewmodel.converter.CountStringConverter.INSTANCE.toCount(callbackArg_0))));
                }
            }
        }
    };
    private androidx.databinding.InverseBindingListener mboundView4androidTextAttrChanged = new androidx.databinding.InverseBindingListener() {
        @Override
        public void onChange() {
            // Inverse of MinimumStringConverter.INSTANCE.toStr(generateInfo.minimum)
            //         is generateInfo.setMinimum((int) MinimumStringConverter.INSTANCE.toMinimum(callbackArg_0))
            java.lang.String callbackArg_0 = androidx.databinding.adapters.TextViewBindingAdapter.getTextString(mboundView4);
            // localize variables for thread safety
            // MinimumStringConverter.INSTANCE.toStr(generateInfo.minimum)
            java.lang.String minimumStringConverterINSTANCEToStrGenerateInfoMinimum = null;
            // generateInfo != null
            boolean generateInfoJavaLangObjectNull = false;
            // generateInfo
            org.csystem.android.app.generator.random.viewmodel.data.GenerateInfo generateInfo = mGenerateInfo;
            // generateInfo.minimum
            int generateInfoMinimum = 0;



            generateInfoJavaLangObjectNull = (generateInfo) != (null);
            if (generateInfoJavaLangObjectNull) {



                if ((org.csystem.android.app.generator.random.viewmodel.converter.MinimumStringConverter.INSTANCE) != (null)) {



                    org.csystem.android.app.generator.random.viewmodel.converter.MinimumStringConverter.INSTANCE.toMinimum(callbackArg_0);


                    generateInfo.setMinimum(((int) (org.csystem.android.app.generator.random.viewmodel.converter.MinimumStringConverter.INSTANCE.toMinimum(callbackArg_0))));
                }
            }
        }
    };
    private androidx.databinding.InverseBindingListener mboundView5androidTextAttrChanged = new androidx.databinding.InverseBindingListener() {
        @Override
        public void onChange() {
            // Inverse of MaximumStringConverter.INSTANCE.toStr(generateInfo.maximum)
            //         is generateInfo.setMaximum((int) MaximumStringConverter.INSTANCE.toMaximum(callbackArg_0))
            java.lang.String callbackArg_0 = androidx.databinding.adapters.TextViewBindingAdapter.getTextString(mboundView5);
            // localize variables for thread safety
            // generateInfo.maximum
            int generateInfoMaximum = 0;
            // MaximumStringConverter.INSTANCE.toStr(generateInfo.maximum)
            java.lang.String maximumStringConverterINSTANCEToStrGenerateInfoMaximum = null;
            // generateInfo != null
            boolean generateInfoJavaLangObjectNull = false;
            // generateInfo
            org.csystem.android.app.generator.random.viewmodel.data.GenerateInfo generateInfo = mGenerateInfo;



            generateInfoJavaLangObjectNull = (generateInfo) != (null);
            if (generateInfoJavaLangObjectNull) {



                if ((org.csystem.android.app.generator.random.viewmodel.converter.MaximumStringConverter.INSTANCE) != (null)) {



                    org.csystem.android.app.generator.random.viewmodel.converter.MaximumStringConverter.INSTANCE.toMaximum(callbackArg_0);


                    generateInfo.setMaximum(((int) (org.csystem.android.app.generator.random.viewmodel.converter.MaximumStringConverter.INSTANCE.toMaximum(callbackArg_0))));
                }
            }
        }
    };

    public ActivityMainBindingImpl(@Nullable androidx.databinding.DataBindingComponent bindingComponent, @NonNull View root) {
        this(bindingComponent, root, mapBindings(bindingComponent, root, 11, sIncludes, sViewsWithIds));
    }
    private ActivityMainBindingImpl(androidx.databinding.DataBindingComponent bindingComponent, View root, Object[] bindings) {
        super(bindingComponent, root, 0
            );
        this.mboundView0 = (android.widget.LinearLayout) bindings[0];
        this.mboundView0.setTag(null);
        this.mboundView1 = (android.widget.EditText) bindings[1];
        this.mboundView1.setTag(null);
        this.mboundView10 = (android.widget.ListView) bindings[10];
        this.mboundView10.setTag(null);
        this.mboundView2 = (android.widget.EditText) bindings[2];
        this.mboundView2.setTag(null);
        this.mboundView3 = (android.widget.EditText) bindings[3];
        this.mboundView3.setTag(null);
        this.mboundView4 = (android.widget.EditText) bindings[4];
        this.mboundView4.setTag(null);
        this.mboundView5 = (android.widget.EditText) bindings[5];
        this.mboundView5.setTag(null);
        this.mboundView6 = (android.widget.Button) bindings[6];
        this.mboundView6.setTag(null);
        this.mboundView7 = (android.widget.Button) bindings[7];
        this.mboundView7.setTag(null);
        this.mboundView8 = (android.widget.Button) bindings[8];
        this.mboundView8.setTag(null);
        this.mboundView9 = (android.widget.Button) bindings[9];
        this.mboundView9.setTag(null);
        setRootTag(root);
        // listeners
        mCallback4 = new org.csystem.android.app.generator.random.generated.callback.OnClickListener(this, 4);
        mCallback2 = new org.csystem.android.app.generator.random.generated.callback.OnClickListener(this, 2);
        mCallback3 = new org.csystem.android.app.generator.random.generated.callback.OnClickListener(this, 3);
        mCallback1 = new org.csystem.android.app.generator.random.generated.callback.OnClickListener(this, 1);
        invalidateAll();
    }

    @Override
    public void invalidateAll() {
        synchronized(this) {
                mDirtyFlags = 0x20L;
        }
        requestRebind();
    }

    @Override
    public boolean hasPendingBindings() {
        synchronized(this) {
            if (mDirtyFlags != 0) {
                return true;
            }
        }
        return false;
    }

    @Override
    public boolean setVariable(int variableId, @Nullable Object variable)  {
        boolean variableSet = true;
        if (BR.serverInfo == variableId) {
            setServerInfo((org.csystem.android.app.generator.random.viewmodel.data.ServerInfo) variable);
        }
        else if (BR.generateInfo == variableId) {
            setGenerateInfo((org.csystem.android.app.generator.random.viewmodel.data.GenerateInfo) variable);
        }
        else if (BR.adapter == variableId) {
            setAdapter((android.widget.ArrayAdapter<java.lang.String>) variable);
        }
        else if (BR.viewModel == variableId) {
            setViewModel((org.csystem.android.app.generator.random.viewmodel.listener.MainActivityListenerViewModel) variable);
        }
        else if (BR.enabled == variableId) {
            setEnabled((boolean) variable);
        }
        else {
            variableSet = false;
        }
            return variableSet;
    }

    public void setServerInfo(@Nullable org.csystem.android.app.generator.random.viewmodel.data.ServerInfo ServerInfo) {
        this.mServerInfo = ServerInfo;
        synchronized(this) {
            mDirtyFlags |= 0x1L;
        }
        notifyPropertyChanged(BR.serverInfo);
        super.requestRebind();
    }
    public void setGenerateInfo(@Nullable org.csystem.android.app.generator.random.viewmodel.data.GenerateInfo GenerateInfo) {
        this.mGenerateInfo = GenerateInfo;
        synchronized(this) {
            mDirtyFlags |= 0x2L;
        }
        notifyPropertyChanged(BR.generateInfo);
        super.requestRebind();
    }
    public void setAdapter(@Nullable android.widget.ArrayAdapter<java.lang.String> Adapter) {
        this.mAdapter = Adapter;
        synchronized(this) {
            mDirtyFlags |= 0x4L;
        }
        notifyPropertyChanged(BR.adapter);
        super.requestRebind();
    }
    public void setViewModel(@Nullable org.csystem.android.app.generator.random.viewmodel.listener.MainActivityListenerViewModel ViewModel) {
        this.mViewModel = ViewModel;
        synchronized(this) {
            mDirtyFlags |= 0x8L;
        }
        notifyPropertyChanged(BR.viewModel);
        super.requestRebind();
    }
    public void setEnabled(boolean Enabled) {
        this.mEnabled = Enabled;
        synchronized(this) {
            mDirtyFlags |= 0x10L;
        }
        notifyPropertyChanged(BR.enabled);
        super.requestRebind();
    }

    @Override
    protected boolean onFieldChange(int localFieldId, Object object, int fieldId) {
        switch (localFieldId) {
        }
        return false;
    }

    @Override
    protected void executeBindings() {
        long dirtyFlags = 0;
        synchronized(this) {
            dirtyFlags = mDirtyFlags;
            mDirtyFlags = 0;
        }
        java.lang.String maximumStringConverterINSTANCEToStrGenerateInfoMaximum = null;
        org.csystem.android.app.generator.random.viewmodel.data.ServerInfo serverInfo = mServerInfo;
        java.lang.String countStringConverterINSTANCEToStrGenerateInfoCount = null;
        java.lang.String portStringConverterINSTANCEToStrServerInfoPort = null;
        int serverInfoPort = 0;
        org.csystem.android.app.generator.random.viewmodel.data.GenerateInfo generateInfo = mGenerateInfo;
        int generateInfoMinimum = 0;
        java.lang.String minimumStringConverterINSTANCEToStrGenerateInfoMinimum = null;
        java.lang.String serverInfoHost = null;
        int generateInfoMaximum = 0;
        android.widget.ArrayAdapter<java.lang.String> adapter = mAdapter;
        int generateInfoCount = 0;
        org.csystem.android.app.generator.random.viewmodel.listener.MainActivityListenerViewModel viewModel = mViewModel;
        boolean enabled = mEnabled;

        if ((dirtyFlags & 0x21L) != 0) {



                if (serverInfo != null) {
                    // read serverInfo.port
                    serverInfoPort = serverInfo.getPort();
                    // read serverInfo.host
                    serverInfoHost = serverInfo.getHost();
                }


                // read PortStringConverter.INSTANCE.toStr(serverInfo.port)
                portStringConverterINSTANCEToStrServerInfoPort = org.csystem.android.app.generator.random.viewmodel.converter.PortStringConverter.INSTANCE.toStr(serverInfoPort);
        }
        if ((dirtyFlags & 0x22L) != 0) {



                if (generateInfo != null) {
                    // read generateInfo.minimum
                    generateInfoMinimum = generateInfo.getMinimum();
                    // read generateInfo.maximum
                    generateInfoMaximum = generateInfo.getMaximum();
                    // read generateInfo.count
                    generateInfoCount = generateInfo.getCount();
                }


                // read MinimumStringConverter.INSTANCE.toStr(generateInfo.minimum)
                minimumStringConverterINSTANCEToStrGenerateInfoMinimum = org.csystem.android.app.generator.random.viewmodel.converter.MinimumStringConverter.INSTANCE.toStr(generateInfoMinimum);
                // read MaximumStringConverter.INSTANCE.toStr(generateInfo.maximum)
                maximumStringConverterINSTANCEToStrGenerateInfoMaximum = org.csystem.android.app.generator.random.viewmodel.converter.MaximumStringConverter.INSTANCE.toStr(generateInfoMaximum);
                // read CountStringConverter.INSTANCE.toStr(generateInfo.count)
                countStringConverterINSTANCEToStrGenerateInfoCount = org.csystem.android.app.generator.random.viewmodel.converter.CountStringConverter.INSTANCE.toStr(generateInfoCount);
        }
        if ((dirtyFlags & 0x24L) != 0) {
        }
        if ((dirtyFlags & 0x30L) != 0) {
        }
        // batch finished
        if ((dirtyFlags & 0x21L) != 0) {
            // api target 1

            androidx.databinding.adapters.TextViewBindingAdapter.setText(this.mboundView1, serverInfoHost);
            androidx.databinding.adapters.TextViewBindingAdapter.setText(this.mboundView2, portStringConverterINSTANCEToStrServerInfoPort);
        }
        if ((dirtyFlags & 0x20L) != 0) {
            // api target 1

            androidx.databinding.adapters.TextViewBindingAdapter.setTextWatcher(this.mboundView1, (androidx.databinding.adapters.TextViewBindingAdapter.BeforeTextChanged)null, (androidx.databinding.adapters.TextViewBindingAdapter.OnTextChanged)null, (androidx.databinding.adapters.TextViewBindingAdapter.AfterTextChanged)null, mboundView1androidTextAttrChanged);
            androidx.databinding.adapters.TextViewBindingAdapter.setTextWatcher(this.mboundView2, (androidx.databinding.adapters.TextViewBindingAdapter.BeforeTextChanged)null, (androidx.databinding.adapters.TextViewBindingAdapter.OnTextChanged)null, (androidx.databinding.adapters.TextViewBindingAdapter.AfterTextChanged)null, mboundView2androidTextAttrChanged);
            androidx.databinding.adapters.TextViewBindingAdapter.setTextWatcher(this.mboundView3, (androidx.databinding.adapters.TextViewBindingAdapter.BeforeTextChanged)null, (androidx.databinding.adapters.TextViewBindingAdapter.OnTextChanged)null, (androidx.databinding.adapters.TextViewBindingAdapter.AfterTextChanged)null, mboundView3androidTextAttrChanged);
            androidx.databinding.adapters.TextViewBindingAdapter.setTextWatcher(this.mboundView4, (androidx.databinding.adapters.TextViewBindingAdapter.BeforeTextChanged)null, (androidx.databinding.adapters.TextViewBindingAdapter.OnTextChanged)null, (androidx.databinding.adapters.TextViewBindingAdapter.AfterTextChanged)null, mboundView4androidTextAttrChanged);
            androidx.databinding.adapters.TextViewBindingAdapter.setTextWatcher(this.mboundView5, (androidx.databinding.adapters.TextViewBindingAdapter.BeforeTextChanged)null, (androidx.databinding.adapters.TextViewBindingAdapter.OnTextChanged)null, (androidx.databinding.adapters.TextViewBindingAdapter.AfterTextChanged)null, mboundView5androidTextAttrChanged);
            this.mboundView6.setOnClickListener(mCallback1);
            this.mboundView7.setOnClickListener(mCallback2);
            this.mboundView8.setOnClickListener(mCallback3);
            this.mboundView9.setOnClickListener(mCallback4);
        }
        if ((dirtyFlags & 0x24L) != 0) {
            // api target 1

            this.mboundView10.setAdapter(adapter);
        }
        if ((dirtyFlags & 0x22L) != 0) {
            // api target 1

            androidx.databinding.adapters.TextViewBindingAdapter.setText(this.mboundView3, countStringConverterINSTANCEToStrGenerateInfoCount);
            androidx.databinding.adapters.TextViewBindingAdapter.setText(this.mboundView4, minimumStringConverterINSTANCEToStrGenerateInfoMinimum);
            androidx.databinding.adapters.TextViewBindingAdapter.setText(this.mboundView5, maximumStringConverterINSTANCEToStrGenerateInfoMaximum);
        }
        if ((dirtyFlags & 0x30L) != 0) {
            // api target 1

            this.mboundView6.setEnabled(enabled);
            this.mboundView7.setEnabled(enabled);
            this.mboundView8.setEnabled(enabled);
            this.mboundView9.setEnabled(enabled);
        }
    }
    // Listener Stub Implementations
    // callback impls
    public final void _internalCallbackOnClick(int sourceId , android.view.View callbackArg_0) {
        switch(sourceId) {
            case 4: {
                // localize variables for thread safety
                // viewModel.handleClearButton()
                kotlin.Unit viewModelHandleClearButton = null;
                // viewModel
                org.csystem.android.app.generator.random.viewmodel.listener.MainActivityListenerViewModel viewModel = mViewModel;
                // viewModel != null
                boolean viewModelJavaLangObjectNull = false;



                viewModelJavaLangObjectNull = (viewModel) != (null);
                if (viewModelJavaLangObjectNull) {


                    viewModelHandleClearButton = viewModel.handleClearButton();
                }
                break;
            }
            case 2: {
                // localize variables for thread safety
                // viewModel.handleConfigurationButton()
                kotlin.Unit viewModelHandleConfigurationButton = null;
                // viewModel
                org.csystem.android.app.generator.random.viewmodel.listener.MainActivityListenerViewModel viewModel = mViewModel;
                // viewModel != null
                boolean viewModelJavaLangObjectNull = false;



                viewModelJavaLangObjectNull = (viewModel) != (null);
                if (viewModelJavaLangObjectNull) {


                    viewModelHandleConfigurationButton = viewModel.handleConfigurationButton();
                }
                break;
            }
            case 3: {
                // localize variables for thread safety
                // viewModel
                org.csystem.android.app.generator.random.viewmodel.listener.MainActivityListenerViewModel viewModel = mViewModel;
                // viewModel != null
                boolean viewModelJavaLangObjectNull = false;
                // viewModel.handleSaveButton()
                kotlin.Unit viewModelHandleSaveButton = null;



                viewModelJavaLangObjectNull = (viewModel) != (null);
                if (viewModelJavaLangObjectNull) {


                    viewModelHandleSaveButton = viewModel.handleSaveButton();
                }
                break;
            }
            case 1: {
                // localize variables for thread safety
                // viewModel.handleGenerateButton()
                kotlin.Unit viewModelHandleGenerateButton = null;
                // viewModel
                org.csystem.android.app.generator.random.viewmodel.listener.MainActivityListenerViewModel viewModel = mViewModel;
                // viewModel != null
                boolean viewModelJavaLangObjectNull = false;



                viewModelJavaLangObjectNull = (viewModel) != (null);
                if (viewModelJavaLangObjectNull) {


                    viewModelHandleGenerateButton = viewModel.handleGenerateButton();
                }
                break;
            }
        }
    }
    // dirty flag
    private  long mDirtyFlags = 0xffffffffffffffffL;
    /* flag mapping
        flag 0 (0x1L): serverInfo
        flag 1 (0x2L): generateInfo
        flag 2 (0x3L): adapter
        flag 3 (0x4L): viewModel
        flag 4 (0x5L): enabled
        flag 5 (0x6L): null
    flag mapping end*/
    //end
}